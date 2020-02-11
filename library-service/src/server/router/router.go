package router

import (
	"github.com/go-chi/chi"
	"library-service/server/handler/logger"
	"library-service/server/handler/response"
	"library-service/server/router/interceptor"
)

func New(a *response.RespHandle) *chi.Mux {
	l := a.Logger()

	r := chi.NewRouter()
	r.Method("GET", "/", logger.NewHandler(a.HandleIndex, l))

	r.Get("/health/liveness", response.HandleLive)
	r.Method("GET", "/health/readiness", logger.NewHandler(a.HandleReady, l))

	// Routes for APIs
	r.Route("/api/v1/library", func(r chi.Router) {
		r.Use(interceptor.ContentTypeJson)

		// Routes for books
		r.Method("GET", "/books", logger.NewHandler(a.HandleListBooks, l))
		r.Method("POST", "/books", logger.NewHandler(a.HandleCreateBook, l))
		r.Method("GET", "/books/{id}", logger.NewHandler(a.HandleReadBook, l))
		r.Method("PUT", "/books/{id}", logger.NewHandler(a.HandleUpdateBook, l))
		r.Method("DELETE", "/books/{id}", logger.NewHandler(a.HandleDeleteBook, l))
	})

	return r
}
