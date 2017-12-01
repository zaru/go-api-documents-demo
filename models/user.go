package user

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type (
	UserModelImpl interface {
		FindAll() []User
	}

	User struct {
		ID   int    `json:"id" db:"id"`
		Name string `json:"name" db:"name"`
		Rank int    `json:"rank" db:"rank"`
		// Tags []Tag  `json:"tags"`
	}

	UserModel struct {
		db *sqlx.DB
	}
)

func NewUserModel(db *sqlx.DB) *UserModel {
	return &UserModel{
		db: db,
	}
}

func (u *UserModel) FindAll() []User {
	users := []User{}
	u.db.Select(&users, "SELECT * FROM user order by id asc")
	return users
}
