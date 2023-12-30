package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *Application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/randomfact", app.RandomFact)
	mux.Get("/fact/{id}", app.OneFact)
	mux.Get("/allfacts", app.AllFacts)


	return mux
}
