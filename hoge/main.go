package main

import "fmt"

type User struct {
	ID   int   `json:"id" db:"id"`
	Tags []Tag `json:"tags"`
}

type Tag struct {
	Name string `json:"name" db:"name"`
}

func main() {
	users := map[int]User{}

	user := User{
		ID: 1,
	}

	users[1] = user

	tag := Tag{
		Name: "hoge",
	}
	user.Tags = append(user.Tags, tag)
	users[1] = user
	// user.Tags = []Tag{}

	fmt.Printf("%v\n", user)
}
