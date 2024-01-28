package models

import (
	"errors"

	"github.com/pedro-hca/go-studies/09-rest-api/db"
	"github.com/pedro-hca/go-studies/09-rest-api/utils"
)

type Users struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u Users) Save() error {
	query := `INSERT INTO users(email, password)
			  VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = userId
	return err
}

func (u *Users) ValidateCredentials() error {
	query := `SELECT id
					,password
			  FROM users
			  WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)
	var retriviedPassword string
	err := row.Scan(&u.ID, &retriviedPassword)
	if err != nil {
		return errors.New("Credentials invalid")
	}
	passwordIsValid := utils.CheckPasswordHash(retriviedPassword, u.Password)
	if !passwordIsValid {
		return errors.New("Credentials invalid")
	}
	return nil
}
