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

// SelectUserID todo doc
func (o Repository) SelectUserID(token string) string {
	var key = ""
	var result string

	rows, err := o.Conn.Query("SELECT user_id FROM authtoken_token WHERE key = $1 LIMIT 1", token)

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		if err := rows.Scan(&key); err != nil {
			log.Fatal(err)
		}
		result = key
	}

	return result
}

// SelectUserPortinformer todo doc
func (o Repository) SelectUserPortinformer(userID string) string {
	fkPortinformer := ""
	result := ""

	rows, err := o.Conn.Query("SELECT fk_portinformer FROM portinformer_managers WHERE fk_user = $1 LIMIT 1", userID)

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		if err := rows.Scan(&fkPortinformer); err != nil {
			log.Println(err)
		}
		result = fkPortinformer
	}
	return result
}

// Close connection
func (o Repository) Close() {
	o.Conn.Close()
}
