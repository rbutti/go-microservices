package handler

import (
	"io"
	"io/ioutil"
	"library-service/util/logger"
	"net"
	"net/http"
	"time"
)

type HandlerWrapper struct {
	handler http.Handler
	logger  *logger.Logger
}

func NewHandler(h http.HandlerFunc, l *logger.Logger) *HandlerWrapper {
	return &HandlerWrapper{
		handler: h,
		logger:  l,
	}
}

func (h *HandlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	le := &logger.LogEntry{
		ReceivedTime:      start,
		RequestMethod:     r.Method,
		RequestURL:        r.URL.String(),
		RequestHeaderSize: logger.HeaderSize(r.Header),
		UserAgent:         r.UserAgent(),
		Referer:           r.Referer(),
		Proto:             r.Proto,
		RemoteIP:          logger.IpFromHostPort(r.RemoteAddr),
	}

	if addr, ok := r.Context().Value(http.LocalAddrContextKey).(net.Addr); ok {
		le.ServerIP = logger.IpFromHostPort(addr.String())
	}
	r2 := new(http.Request)
	*r2 = *r
	rcc := &logger.ReadCounterCloser{R: r.Body}
	r2.Body = rcc
	W2 := &logger.ResponseStats{W: w}

	h.handler.ServeHTTP(W2, r2)

	le.Latency = time.Since(start)
	if rcc.ERR == nil && rcc.R != nil {
		// If the handler hasn't encountered an error in the Body (like EOF),
		// then consume the rest of the Body to provide an accurate rcc.n.
		io.Copy(ioutil.Discard, rcc)
	}
	le.RequestBodySize = rcc.N
	le.Status = W2.CODE
	if le.Status == 0 {
		le.Status = http.StatusOK
	}
	le.ResponseHeaderSize, le.ResponseBodySize = W2.Size()
	h.logger.Info().
		Time("received_time", le.ReceivedTime).
		Str("method", le.RequestMethod).
		Str("url", le.RequestURL).
		Int64("header_size", le.RequestHeaderSize).
		Int64("body_size", le.RequestBodySize).
		Str("agent", le.UserAgent).
		Str("referer", le.Referer).
		Str("proto", le.Proto).
		Str("remote_ip", le.RemoteIP).
		Str("server_ip", le.ServerIP).
		Int("status", le.Status).
		Int64("resp_header_size", le.ResponseHeaderSize).
		Int64("resp_body_size", le.ResponseBodySize).
		Dur("latency", le.Latency).
		Msg("")
}
