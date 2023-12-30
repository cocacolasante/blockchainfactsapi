package main

import (
	"flag"
	"log"

	"github.com/cocacolasante/blockchainfacts/database/postgresrepo"
)

const PORT = 8080


func main() {

	var app Application

	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=blockchain facts sslmode=disable timezone=utc connect_timeout=5", "Postgres connection string")
	flag.StringVar(&app.Domain, "domain", "example.com", "domain")

	flag.Parse()

	app.DSN = "host=localhost port=5432 user=postgres password=postgres dbname=blockchainfacts sslmode=disable timezone=utc connect_timeout=5"

	conn, err := app.ConnectToDb()
	if err != nil {
		log.Println("main")
		log.Fatal(err)
	}

	app.DB = &postgresrepo.PostgresRepo{DB: conn}

	defer app.DB.Connection().Close()

	oneFact := app.DB.OneFact()
	log.Println(oneFact)

}
