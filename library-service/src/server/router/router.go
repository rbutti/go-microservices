package router

import (
	"github.com/go-chi/chi"
	"library-service/server/handler"
	"library-service/server/router/interceptor"
)

func New(a *handler.Handler) *chi.Mux {
	l := a.Logger()

	r := chi.NewRouter()
	r.Method("GET", "/", handler.NewHandler(a.HandleIndex, l))

	r.Get("/health/liveness", handler.HandleLive)
	r.Method("GET", "/health/readiness", handler.NewHandler(a.HandleReady, l))

	// Routes for APIs
	r.Route("/api/v1/library", func(r chi.Router) {
		r.Use(interceptor.ContentTypeJson)

		// Routes for books
		r.Method("GET", "/books", handler.NewHandler(a.HandleListBooks, l))
		r.Method("POST", "/books", handler.NewHandler(a.HandleCreateBook, l))
		r.Method("GET", "/books/{id}", handler.NewHandler(a.HandleReadBook, l))
		r.Method("PUT", "/books/{id}", handler.NewHandler(a.HandleUpdateBook, l))
		r.Method("DELETE", "/books/{id}", handler.NewHandler(a.HandleDeleteBook, l))
	})

	return r
}
