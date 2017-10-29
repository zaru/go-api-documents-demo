package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/users", getUser)
	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c echo.Context) error {
	u := User{
		Name: "zaru",
	}
	return c.JSON(http.StatusOK, u)
}

type User struct {
	Name string `json:"name"`
}
