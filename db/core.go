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
	conn *sql.DB
}

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

// SelectUserID todo doc
func (o Repository) SelectUserID(token string) []string {
	var key = ""
	var result = []string{}

	rows, err := o.conn.Query("SELECT user_id FROM authtoken_token WHERE key = $1", token)

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		if err := rows.Scan(&key); err != nil {
			log.Fatal(err)
		}
		result = append(result, key)
	}

	return result
}
