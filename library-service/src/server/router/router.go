package router

import (
	"github.com/go-chi/chi"
	"library-service/server/handler/logger"
	"library-service/server/handler/response"
)

func New(a *response.RespHandle) *chi.Mux {
	l := a.Logger()

	r := chi.NewRouter()
	r.Method("GET", "/", logger.NewHandler(a.HandleIndex, l))

	return r
}
