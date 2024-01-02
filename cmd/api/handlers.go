package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/cocacolasante/blockchainfacts/models"
	"github.com/go-chi/chi/v5"
)

type JSONPayload struct {
	ID   int    `json:"fact_id"`
	Fact string `json:"fact_text"`
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


func (app *Application) AddFact(w http.ResponseWriter, r *http.Request) {
	var readfact *models.BCFact

	
	defer r.Body.Close()
	
	err := app.ReadJSONFromBody(r, &readfact)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var payload = struct {
			HasError bool   `json:"has_error"`
			Error    string `json:"error_message"`
		}{
			HasError: true,
			Error:    err.Error(),
		}

		out, err := json.Marshal(payload)
		if err != nil {
			log.Println("error marshalling json")
			return
		}
		w.Write(out)
		return
	}

	if readfact.Fact == ""{
		w.WriteHeader(http.StatusBadRequest)
		var payload = struct {
			HasError bool   `json:"has_error"`
			Error    string `json:"error_message"`
		}{
			HasError: true,
			Error:    "Empty Fact",
		}

		out, err := json.Marshal(payload)
		if err != nil {
			log.Println("error marshalling json")
			return
		}
		w.Write(out)
		return
	}
	fact, err := app.DB.AddFact(readfact.Fact)
	if err != nil {
		
		w.WriteHeader(http.StatusBadRequest)
		var payload = struct {
			HasError bool   `json:"has_error"`
			Error    string `json:"error_message"`
		}{
			HasError: true,
			Error:    err.Error(),
		}

		out, err := json.Marshal(payload)
		if err != nil {
			log.Println("error marshalling json")
			return
		}
		w.Write(out)
		return
	}

	out, err := json.Marshal(fact)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var payload = struct {
			HasError bool   `json:"has_error"`
			Error    string `json:"error_message"`
		}{
			HasError: true,
			Error:    "Error marshalling data",
		}

		out, err := json.Marshal(payload)
		if err != nil {
			log.Println("error marshalling json")
			return
		}
		w.Write(out)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(out)

}

func (app *Application) DeleteFact(w http.ResponseWriter, r *http.Request){
	var readfact *models.BCFact

	
	defer r.Body.Close()
	
	err := app.ReadJSONFromBody(r, &readfact)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var payload = struct {
			HasError bool   `json:"has_error"`
			Error    string `json:"error_message"`
		}{
			HasError: true,
			Error:    err.Error(),
		}

		out, err := json.Marshal(payload)
		if err != nil {
			log.Println("error marshalling json")
			return
		}
		w.Write(out)
		return
	}
	

	if readfact.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		var payload = struct {
			HasError bool   `json:"has_error"`
			Error    string `json:"error_message"`
		}{
			HasError: true,
			Error:    "No fact number selected",
		}

		out, err := json.Marshal(payload)
		if err != nil {
			log.Println("error marshalling json")
			return
		}
		w.Write(out)
		return
	}

	deleted, err := app.DB.DeleteFact(readfact.ID)
	if err !=nil || !deleted {
		w.WriteHeader(http.StatusBadRequest)
		var payload = struct {
			HasError bool   `json:"has_error"`
			Error    string `json:"error_message"`
		}{
			HasError: true,
			Error:    "Fact was not deleted",
		}

		out, err := json.Marshal(payload)
		if err != nil {
			log.Println("error marshalling json")
			return
		}
		w.Write(out)
		return
	}

	var payload = struct {
		Deleted bool `json:"deleted"`
		
	}{
		Deleted: true,
	}
	out, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var payload = struct {
			HasError bool   `json:"has_error"`
			Error    string `json:"error_message"`
		}{
			HasError: true,
			Error:    err.Error(),
		}

		out, err := json.Marshal(payload)
		if err != nil {
			log.Println("error marshalling json")
			return
		}
		w.Write(out)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(out)
}