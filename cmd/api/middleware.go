package main

import "net/http"

func (app *Application) EnableCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the allowed origin
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Methods", "GET")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, X-CSRF-Token, Authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true") // Set to 'true' for credentials support
			w.WriteHeader(http.StatusOK)
			return
		}

		// Allow credentials in the actual response
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Continue with the request handling
		h.ServeHTTP(w, r)
	})
}
