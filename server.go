package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"gopkg.in/boj/redistore.v1"
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

	store, err := redistore.NewRediStore(10, "tcp", "kvs:6379", "", []byte("secret-key"))
	ss := sessions.Store(store)

	e := echo.New()
	e.Use(session.Middleware(aa))

	e.GET("/users", getUser)
	e.GET("/me", me)
	e.Logger.Fatal(e.Start(":1323"))
}

func me(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["foo"] = "bar"
	sess.Save(c.Request(), c.Response())
	return c.NoContent(http.StatusOK)
}

func getUser(c echo.Context) error {

	db, err := sqlx.Connect("mysql", "root:@("+os.Getenv("DB_HOST")+":3306)/"+os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatalln(err)
	}

	rows, _ := db.Queryx("SELECT * FROM user")
	users := map[int]User{}
	for rows.Next() {
		user := User{}
		rows.StructScan(&user)
		users[user.ID] = user
	}

	uIDs := []int{}
	for _, v := range users {
		uIDs = append(uIDs, v.ID)
	}

	query, args, _ := sqlx.In("SELECT user_id, name FROM tag WHERE user_id IN (?);", uIDs)
	query = db.Rebind(query)
	tags, _ := db.Queryx(query, args...)
	for tags.Next() {
		tag := Tag{}
		tags.StructScan(&tag)
		user := users[tag.UserID]
		user.Tags = append(user.Tags, tag)
		users[tag.UserID] = user
	}

	output := []User{}
	for _, v := range users {
		output = append(output, v)
	}

	return c.JSON(http.StatusOK, output)
}

// User struct
type User struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Rank int    `json:"rank" db:"rank"`
	Tags []Tag  `json:"tags"`
}

// Tag struct
type Tag struct {
	UserID int    `json:"-" db:"user_id"`
	Name   string `json:"name" db:"name"`
}
