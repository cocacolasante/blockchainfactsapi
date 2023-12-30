package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var JSONPayload struct {
	ID   int    `json:"fact_id"`
	Fact string `json:"fact"`
}

func (app *Application) RandomFact(w http.ResponseWriter, r *http.Request) {
	fact, err := app.DB.OneFact()
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
