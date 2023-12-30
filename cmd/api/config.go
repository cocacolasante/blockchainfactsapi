package main

import (
	"database/sql"
	"log"

	"github.com/cocacolasante/blockchainfacts/database/postgresrepo"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Application struct {
	Domain string
	DSN    string
	Port   int
	DB     *postgresrepo.PostgresRepo
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return db, nil
}

func (app *Application) ConnectToDb() (*sql.DB, error) {
	conn, err := openDB(app.DSN)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("Connected to postgres")

	return conn, nil
}
