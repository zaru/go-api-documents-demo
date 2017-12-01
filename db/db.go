package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

func DBConnect() *sqlx.DB {
	db, err := sqlx.Connect("mysql", "root:@("+os.Getenv("DB_HOST")+":3306)/"+os.Getenv("DB_NAME"))
	if err != nil && os.Getenv("APP_ENV") != "test" {
		log.Fatalln(err)
	}
	return db
}
