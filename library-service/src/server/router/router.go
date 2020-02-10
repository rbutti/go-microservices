package router

import (
	"github.com/go-chi/chi"
	"library-service/server/handler"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.MethodFunc("GET", "/", handler.HandleIndex)

	return r
}
