package router

import (
	"github.com/go-chi/chi"

	"go-world/app"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.MethodFunc("GET", "/", app.HandleIndex)

	return r
}