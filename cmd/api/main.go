package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cocacolasante/blockchainfacts/database/postgresrepo"
)

const PORT = 443

func main() {

	certFile := "/etc/letsencrypt/live/api.anthonycolasante.com/fullchain.pem"
	keyFile := "/etc/letsencrypt/live/api.anthonycolasante.com/privkey.pem"

	var app Application

	app.DSN = "host=localhost port=5432 user=admin password=adminpass dbname=blockchainfacts sslmode=disable timezone=utc connect_timeout=5"
	app.Port = PORT
	conn, err := app.ConnectToDb()
	if err != nil {

		log.Fatal(err)
	}
	defer conn.Close()
	app.DB = &postgresrepo.PostgresRepo{DB: conn}

	// defer app.DB.Connection().Close()

	pop, err := app.CheckIfDBIsPopulated()
	if err != nil {
		log.Println(err)
	}

	if pop {

		app.createTable()
		app.addData()
	}
	// 1) (Optional) Start an HTTP->HTTPS redirect on port 80
	go func() {
		log.Println("Starting HTTP redirect server on port 80...")
		// Any request on port 80 gets redirected to https://...
		err := http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusMovedPermanently)
		}))
		if err != nil {
			log.Printf("HTTP redirect server closed with error: %v", err)
		}
	}()

	// 2) Start HTTPS server on port 443
	log.Println("Starting HTTPS server on port:", app.Port)
	err = http.ListenAndServeTLS(
		fmt.Sprintf(":%d", app.Port),
		certFile, // path to your certificate
		keyFile, // path to your private key
		app.routes(), // your route multiplexer
	)
	if err != nil {
		log.Fatal(err)
	}

}
