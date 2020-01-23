package db

import (
	"github.com/go-pg/pg/v9"
)

// Connector create gorm instance
func Connector() *pg.DB {

	db := pg.Connect(&pg.Options{
		User:     "shipreporting_user",
		Password: "",
		Database: "shipreporting_sw",
	})

	return db
}
