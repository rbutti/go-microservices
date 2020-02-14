package log

import (
	"errors"
	"io"
	"net"
	"net/http"
	"time"
)

type LogEntry struct {
	ReceivedTime      time.Time
	RequestMethod     string
	RequestURL        string
	RequestHeaderSize int64
	RequestBodySize   int64
	UserAgent         string
	Referer           string
	Proto             string

	RemoteIP string
	ServerIP string

	Status             int
	ResponseHeaderSize int64
	ResponseBodySize   int64
	Latency            time.Duration
}

type ReadCounterCloser struct {
	R   io.ReadCloser
	N   int64
	ERR error
}

type WriteCounter int64

type ResponseStats struct {
	W     http.ResponseWriter
	hsize int64
	wc    WriteCounter
	CODE  int
}

func IpFromHostPort(hp string) string {
	h, _, err := net.SplitHostPort(hp)
	if err != nil {
		return ""
	}
	if len(h) > 0 && h[0] == '[' {
		return h[1 : len(h)-1]
	}
	return h
}

func (rcc *ReadCounterCloser) Read(p []byte) (n int, err error) {
	if rcc.ERR != nil {
		return 0, rcc.ERR
	}
	n, rcc.ERR = rcc.R.Read(p)
	rcc.N += int64(n)
	return n, rcc.ERR
}

func (rcc *ReadCounterCloser) Close() error {
	rcc.ERR = errors.New("read from closed reader")
	return rcc.R.Close()
}

func (wc *WriteCounter) Write(p []byte) (n int, err error) {
	*wc += WriteCounter(len(p))
	return len(p), nil
}

func HeaderSize(h http.Header) int64 {
	var wc WriteCounter
	h.Write(&wc)
	return int64(wc) + 2 // for CRLF
}

func (r *ResponseStats) Header() http.Header {
	return r.W.Header()
}

func (r *ResponseStats) WriteHeader(statusCode int) {
	if r.CODE != 0 {
		return
	}
	r.hsize = HeaderSize(r.W.Header())
	r.W.WriteHeader(statusCode)
	r.CODE = statusCode
}

func (r *ResponseStats) Write(p []byte) (n int, err error) {
	if r.CODE == 0 {
		r.WriteHeader(http.StatusOK)
	}
	n, err = r.W.Write(p)
	r.wc.Write(p[:n])
	return
}

func (r *ResponseStats) Size() (hdr, body int64) {
	if r.CODE == 0 {
		return HeaderSize(r.W.Header()), 0
	}
	// Use the header size from the time WriteHeader was called.
	// The Header map can be mutated after the call to add HTTP Trailers,
	// which we don't want to count.
	return r.hsize, int64(r.wc)
}
