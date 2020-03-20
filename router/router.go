package router

import (
	"go-world/app"

	"github.com/go-chi/chi"
)

func New() *chi.Mux {
	router := chi.NewRouter()

	router.MethodFunc("GET", "/", app.HandleIndex)

	return router
}
