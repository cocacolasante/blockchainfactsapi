package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var JSONPayload struct {
	ID   int    `json:"fact_id"`
	Fact string `json:"fact"`
}

func (app *Application) RandomFact(w http.ResponseWriter, r *http.Request) {
	fact, err := app.DB.OneFactRandom()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var payload = struct {
			HasError bool   `json:"has_error"`
			Error    string `json:"error_message"`
		}{
			HasError: true,
			Error:    "No Facts In Database",
		}

		out, err := json.Marshal(payload)
		if err != nil {
			log.Println("error marshalling json")
			return
		}
		w.Write(out)
	}
	w.Header().Add("Content-Type", "application/json")

	out, err := json.Marshal(fact)
	if err != nil {
		log.Println("error marshalling json")
		return
	}

	w.Write(out)

}
func (app *Application) OneFact(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	log.Println(id)
	intid, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var payload = struct {
			HasError bool   `json:"has_error"`
			Error    string `json:"error_message"`
		}{
			HasError: true,
			Error:    "No id given",
		}

		out, err := json.Marshal(payload)
		if err != nil {
			log.Println("error marshalling json")
			return
		}
		w.Write(out)
	}
	fact, err := app.DB.OneFact(intid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var payload = struct {
			HasError bool   `json:"has_error"`
			Error    string `json:"error_message"`
		}{
			HasError: true,
			Error:    "Fact Id not in Database",
		}

		out, err := json.Marshal(payload)
		if err != nil {
			log.Println("error marshalling json")
			return
		}
		w.Write(out)
	}
	w.Header().Add("Content-Type", "application/json")

	out, err := json.Marshal(fact)
	if err != nil {
		log.Println("error marshalling json")
		return
	}

	w.Write(out)

}
func (app *Application) AllFacts(w http.ResponseWriter, r *http.Request) {
	
	facts, err := app.DB.AllFacts()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var payload = struct {
			HasError bool   `json:"has_error"`
			Error    string `json:"error_message"`
		}{
			HasError: true,
			Error:    "Fact Id not in Database",
		}

		out, err := json.Marshal(payload)
		if err != nil {
			log.Println("error marshalling json")
			return
		}
		w.Write(out)
	}
	w.Header().Add("Content-Type", "application/json")

	out, err := json.Marshal(facts)
	if err != nil {
		log.Println("error marshalling json")
		return
	}

	w.Write(out)

}

