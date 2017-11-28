package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	appENV := os.Getenv("APP_ENV")
	if len(appENV) == 0 {
		appENV = "development"
	}

	err := godotenv.Load(".env." + appENV)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()
	e.GET("/users", getUser)
	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c echo.Context) error {

	db, err := sqlx.Connect("mysql", "root:@(localhost:3306)/"+os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatalln(err)
	}
	users := []User{}
	db.Select(&users, "SELECT id, name FROM user")

	for _, v := range users {
		append(v.Tags, "hoge")
	}

	return c.JSON(http.StatusOK, users)
}

type User struct {
	ID   int      `json:"id" db:"id"`
	Name string   `json:"name" db:"name"`
	Rank string   `json:"rank" db:"rank"`
	Tags []string `json:"tags"`
}
