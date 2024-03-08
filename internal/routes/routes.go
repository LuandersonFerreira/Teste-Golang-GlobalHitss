package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetRoutes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	r.Post("/", CreateUser)
	// r.Put("/", UpdateUser)
	// r.Delete("/", DeleteUser)
	// r.Get("/list", ListUser)
	// r.Get("/", GetUser)
}
