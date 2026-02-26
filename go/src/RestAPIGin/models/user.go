package models

import (
	"errors"
	"fmt"

	"example.com/restApiGin/db"
	"example.com/restApiGin/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) SaveUser() error {
	query := `INSERT INTO users(email, password) VALUES(?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashPassword, _ := utils.Hash(user.Password)
	result, err := stmt.Exec(user.Email, hashPassword)
	if err != nil {
		return err
	}
	userID, err := result.LastInsertId()
	user.ID = userID
	return err
}

func (user *User) ValidateUser() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, user.Email)
	var retrievedPassword string
	err := row.Scan(&user.ID, &retrievedPassword)
	if err != nil {
		fmt.Println(err)
		return errors.New("invalid credentials")
	}
	passwordIsValid := utils.CheckPassword(user.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("invalid credentials")
	}
	return nil
}
