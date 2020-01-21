package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/deeper-x/shipreporting-platform/utils"
	"github.com/joho/godotenv"

	// pq lib as a blank import
	_ "github.com/lib/pq"
)

// Repository main db object
type Repository struct {
	Conn *sql.DB
}

type repo Repository

// Connector return db
func Connector() *sql.DB {
	err := godotenv.Load(utils.DotenvFile)

	if err != nil {
		log.Println(err)
	}

	db, err := sql.Open("postgres", os.Getenv("DB_DSN"))

	if err != nil {
		log.Fatal(err)
	}

	return db
}

// NewRepository create connection
func NewRepository(conn *sql.DB) Repository {
	repo := Repository{Conn: conn}
	return repo
}

// Close connection
func (r Repository) Close() {
	r.Conn.Close()
}
