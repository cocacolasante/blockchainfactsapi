package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

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

// returns true if the count is zero
func (app *Application) CheckIfDBIsPopulated() (bool, error) {

	var count int
	query := `SELECT COUNT(*) FROM facts;`
	
	err := app.DB.DB.QueryRow(query).Scan(&count)
	
	if err != nil {
		if err.Error() == "ERROR: relation \"facts\" does not exist (SQLSTATE 42P01)"{
			return true, nil
		}
		return false, err
	}
	
	if count == 0 {
		return true, nil
	}

	
	log.Println("Database is populated")
	// If the count is > 0, we can safely say the DB is populated.
	return false, nil
}
func (app *Application) PopulateDb() error {
	err := app.createTable()
	if err != nil{
		return err
	}

	err = app.addData()
	if err != nil {
		return err
	}
	return nil

}
func (app *Application) createTable() error {
	err := app.DB.CreateTable()
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (app *Application) addData() error {
	log.Println("Running populating data script")
	f, err := os.Open("sql/pop_db.sql")
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()
	
	dataBytes, err := io.ReadAll(f)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	
	statements := strings.Split(string(dataBytes), ";")
	for _, stmt := range statements {
		
		// Trim whitespace/newlines
		cleaned := strings.TrimSpace(stmt)
		if cleaned == "" {
			continue
		}

		// Execute each statement
		_, execErr := app.DB.DB.Exec(cleaned)
		if execErr != nil {
			return fmt.Errorf("error executing statement (%s): %w", cleaned, execErr)
		}
	}

	return nil
}

