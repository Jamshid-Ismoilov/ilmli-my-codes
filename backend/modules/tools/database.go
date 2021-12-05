package tools

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	host     = "localhost"
	user     = "username"
	password = "password"
	dbname   = "ilmli"
	port     = 5432
)

var DB_CONFIG = fmt.Sprintf(
	"host=%s user=%s password=%s dbname=%s port=%d",
	host, user, password, dbname, port,
)

type DB struct {
	Conn *sql.DB
}

func (d DB) ConnectionPQ() *sql.DB {
	db, err := sql.Open("postgres", DB_CONFIG)

	if err != nil {
		log.Printf("failed to connect database: %v", err)
	}

	return db
}
