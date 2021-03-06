package models

import "errors"

var ErrUsernameNotFound = errors.New("person.model: user with given username not found")

type Person struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required" db:"name"`
	Username string `json:"username" binding:"required" db:"username"`
	Password string `json:"password" binding:"required" db:"password_hash"`
}
