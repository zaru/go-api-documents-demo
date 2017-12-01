package user

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/zaru/go-api-documents-demo/db"
)

func TestFindAll(t *testing.T) {

	d := db.DBConnect()
	um := NewUserModel(d)

	u1 := User{
		ID:   1,
		Name: "zaru",
		Rank: 1,
	}

	u := um.FindAll()

	expect := []User{}
	expect = append(expect, u1)
	assert.Equal(t, expect, u)
}

func setup() {
	err := godotenv.Load("../.env.test")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	exec.Command("mysql", "-h", os.Getenv("DB_HOST"), "-u", "root", "-e", "create database "+os.Getenv("DB_NAME")).Run()
	exec.Command("goose", "-dir", "../migrations", "mysql", "root@("+os.Getenv("DB_HOST")+")/"+os.Getenv("DB_NAME"), "up").Run()
	cmd := exec.Command("mysql", "-h", os.Getenv("DB_HOST"), "-u", "root", os.Getenv("DB_NAME"))
	readFile := "../seed/test.sql"
	input, err := ioutil.ReadFile(readFile)
	if err != nil {
		panic(err)
	}
	stdin, _ := cmd.StdinPipe()
	stdin.Write(input)
	stdin.Close()

	cmd.Run()
}

func clear() {
	exec.Command("mysql", "-h", os.Getenv("DB_HOST"), "-u", "root", "-e", "drop database "+os.Getenv("DB_NAME")).Run()
}

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	clear()

	os.Exit(ret)
}
