package models

import (
	"golang.org/x/crypto/bcrypt"

	"go-crud/db"
)

type ResponseRegister struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type UserRegister struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

func CreateUser(username, password, fullName string) error {
	con := db.CreateCon()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	sqlStatement := `INSERT INTO users (username, password, full_name) VALUES (?, ?, ?)`
	_, err = con.Exec(sqlStatement, username, string(hashedPassword), fullName)
	if err != nil {
		return err
	}

	return nil
}
