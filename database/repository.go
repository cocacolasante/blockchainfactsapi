package dbrepo

import "database/sql"

type Database interface {
	Connection() *sql.DB
}
