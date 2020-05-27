package models

import (
	"github.com/loremcookie/go-home/backend/internal/database"
	"github.com/loremcookie/go-home/backend/internal/passhash"
)

//User is the user struct that defines a user.
//The user struct is the user information in the database.
type User struct {
	Username string
	Password string
	Admin    bool
}

//NewUser saves a user in database
func NewUser(user *User) error {
	var err error

	//Put user in database
	err = database.Update("USER", user.Username, user)
	if err != nil {
		return err
	}

	return nil
}

//GetUser retrieves user from database by username
func GetUser(username string) (*User, error) {
	var err error

	//Create empty user struct to store retrieved user in
	var user User

	//Get user from database
	err = database.View("USER", username, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

//DeleteUser deletes user from database by username
func DeleteUser(username string) error {
	var err error

	err = database.Delete("USER", username)
	if err != nil {
		return err
	}

	return nil
}

//ValidUser returns a bool based of a user is valid
func ValidUser(login *Login) bool {
	var err error

	//Checks if data is privet
	if len(login.Username) == 0 && len(login.Password) == 0 {
		return false
	}

	//Get user with matching username
	user, err := GetUser(login.Username)
	if err != nil {
		return false
	}

	//Check if username match
	if login.Username != user.Username {
		return false
	}

	//Check if password matches password hash
	if passhash.MatchString(login.Password, user.Password) {
		return false
	}

	return true
}
