package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (app *Application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(app.EnableCORS)

	mux.Get("/randomfact", app.RandomFact)
	mux.Get("/fact/{id}", app.OneFact)
	mux.Get("/allfacts", app.AllFacts)
	mux.Post("/addfact", app.AddFact)

	mux.Delete("/deletefact", app.DeleteFact)
	mux.Patch("/updatefact", app.UpdateFact)

	return mux
}
