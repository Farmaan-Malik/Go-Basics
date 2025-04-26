package models

import (
	"errors"
	"example/rest-api/db"
	"example/rest-api/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
	INSERT INTO users(email,password) VALUES (?,?)
	`
	sql, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer sql.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	result, err := sql.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.Id = id
	return nil
}

func (u *User) ValidateCredentials() error {
	query := `
SELECT id,password FROM users WHERE email = ?
`
	row := db.DB.QueryRow(query, u.Email)
	var hashedPassword string
	err := row.Scan(&u.Id, &hashedPassword)
	if err != nil {
		return errors.New("invalid Credentials")
	}
	isValidCred := utils.CompareHash(hashedPassword, u.Password)
	if !isValidCred {
		return errors.New("invalid Credentials")
	}
	return nil
}
