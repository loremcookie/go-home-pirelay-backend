package models

import (
	"encoding/json"
	"github.com/loremcookie/go-home-pirelay-backend/internal/database"
	"github.com/loremcookie/go-home-pirelay-backend/internal/passhash"
)

//User is the user struct that defines a user.
//The user struct is the user information in the database.
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Admin    bool   `json:"admin"`
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

//GetAllUser returns all user that are in the database
func GetAllUser() ([]User, error) {
	var err error

	//Get all keys and values of the USER bucket
	keyValues, err := database.GetAll("USER")
	if err != nil {
		return nil, err
	}

	//Create a slice of users to store users in
	var users []User

	//Iterate over the keys and values of the bucket
	for _, val := range keyValues {

		//Iterate over the keys and values of the key an value pair's of the bucket
		for _, value := range val {
			//Create user to decode json to
			var user User

			//Decode the values from json to user
			err := json.Unmarshal(value.([]byte), &user)
			if err != nil {
				return nil, err
			}

			//Append user to list of users
			users = append(users, user)

		}
	}

	//Return all user in the USER bucket
	return users, nil
}

//ValidUser returns a bool based of a user is valid
func ValidUser(login *Login) bool {
	var err error

	//Validate function input
	if login.Username == "" {
		return false
	}
	if login.Password == "" {
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
	if !passhash.MatchString(login.Password, user.Password) {
		return false
	}

	return true
}
