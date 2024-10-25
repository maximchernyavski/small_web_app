package model

import (
	"errors"

	"example.com/web_shit/db"
	"example.com/web_shit/utils"
)

type User struct {
	ID       int64
	Login    string `binding:"required"`
	Password string `binding:"required"`
	IsAdmin  bool   `json:"isAdmin"`
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE login = ?"
	row := db.DB.QueryRow(query, u.Login)

	var retrievedPassword string
	var isAdmin bool
	err := row.Scan(&u.ID, &retrievedPassword, &isAdmin)

	if err != nil {
		return errors.New("credentials invalid")
	}

	if u.IsAdmin != isAdmin {
		return errors.New("credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}

func (u *User) Save() error {
	query := "INSERT INTO users(login, password) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Login, hashedPassword)
	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = userID
	return nil
}
