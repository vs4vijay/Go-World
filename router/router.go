package router

import (
	"github.com/go-chi/chi"
	"go-world/app"
)

func New() *chi.Mux {
	router := chi.NewRouter()

	router.MethodFunc("GET", "/go", app.HandleIndex)

	return router
}
