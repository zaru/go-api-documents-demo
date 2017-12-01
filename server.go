package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/zaru/go-api-documents-demo/db"
	"github.com/zaru/go-api-documents-demo/handlers"
	"github.com/zaru/go-api-documents-demo/models"
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

	d := db.DBConnect()
	h := users.NewHandler(user.NewUserModel(d))

	e.GET("/users", h.GetIndex)

	e.Logger.Fatal(e.Start(":1323"))
}

//
// func getUser(c echo.Context) error {
//
// 	db, err := sqlx.Connect("mysql", "root:@("+os.Getenv("DB_HOST")+":3306)/"+os.Getenv("DB_NAME"))
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
//
// 	rows, _ := db.Queryx("SELECT * FROM user")
// 	users := map[int]User{}
// 	for rows.Next() {
// 		user := User{}
// 		rows.StructScan(&user)
// 		users[user.ID] = user
// 	}
//
// 	uIDs := []int{}
// 	for _, v := range users {
// 		uIDs = append(uIDs, v.ID)
// 	}
//
// 	query, args, _ := sqlx.In("SELECT user_id, name FROM tag WHERE user_id IN (?);", uIDs)
// 	query = db.Rebind(query)
// 	tags, _ := db.Queryx(query, args...)
// 	for tags.Next() {
// 		tag := Tag{}
// 		tags.StructScan(&tag)
// 		user := users[tag.UserID]
// 		user.Tags = append(user.Tags, tag)
// 		users[tag.UserID] = user
// 	}
//
// 	output := []User{}
// 	for _, v := range users {
// 		output = append(output, v)
// 	}
//
// 	return c.JSON(http.StatusOK, output)
// }
//
// // Tag struct
// type Tag struct {
// 	UserID int    `json:"-" db:"user_id"`
// 	Name   string `json:"name" db:"name"`
// }
