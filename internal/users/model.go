package users

import "time"

type UserCreateForm struct {
	Name     string
	Email  string
	PassHash     string
}

type User struct {
	ID        int `db:"id"`
	Name      string `db:"name"`
	Email     string `db:"email"`
	PassHash   string `db:"pass_hash"`
	CreatedAt time.Time `db:"createdat"`
}