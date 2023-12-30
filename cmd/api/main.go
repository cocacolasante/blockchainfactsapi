package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cocacolasante/blockchainfacts/database/postgresrepo"
)

const PORT = 8080

func main() {

	var app Application

	app.DSN = "host=localhost port=5432 user=postgres password=postgres dbname=blockchainfacts sslmode=disable timezone=utc connect_timeout=5"
	app.Port = PORT
	conn, err := app.ConnectToDb()
	if err != nil {
		log.Println("main")
		log.Fatal(err)
	}

	app.DB = &postgresrepo.PostgresRepo{DB: conn}

	defer app.DB.Connection().Close()

	log.Println("Starting application on port:", app.Port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", app.Port), app.routes())
	if err != nil {
		log.Fatal(err)

	}

	

}
