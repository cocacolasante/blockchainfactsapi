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

	log.Println("Starting application on port:", app.Port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", app.Port), app.routes())
	if err != nil {
		log.Fatal(err)

	}

	

}
