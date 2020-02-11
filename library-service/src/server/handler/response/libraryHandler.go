package response

import (
	"net/http"
)

func (app *RespHandle) HandleListBooks(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("[]"))
}

func (app *RespHandle) HandleCreateBook(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}

func (app *RespHandle) HandleReadBook(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}

func (app *RespHandle) HandleUpdateBook(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}

func (app *RespHandle) HandleDeleteBook(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}
