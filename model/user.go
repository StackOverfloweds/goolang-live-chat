package model

import (
	"database/sql"
	"go-live-chat/config"

	"golang.org/x/crypto/bcrypt"
)

//USER Struct
type User struct {
	ID		 	int
	Username 	string
	Email 		string
	Password 	string
}

//register a new user
func (U * User) Register() error {
	hashedPassword, err:= bcrypt.GenerateFromPassword([]byte (U.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = config.DB.Exec("INSERT INTO  users (username, email, password_hash) VALUES (?,?,?)",
	U.Username, U.Email, string(hashedPassword))
	return err
}

//Auth user login
func Authenticate (email, password string) (*User, error) {
	user := &User{}
	row := config.DB.QueryRow("SELECT id, username, password_hash FROM users WHERE email = ?",email)

	if err := row.Scan(&user.ID, &user.Username, &user.Password); err != nil {
		return nil, sql.ErrNoRows
	}	

	return user, nil
}