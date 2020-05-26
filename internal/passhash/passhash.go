/*
This package is used for hashing and salting strings mainly passwords.
This is a modified version of github.com/josephspurrier/gowebapp/tree/master/vendor/app/shared/passhash
*/
package passhash

import (
	"golang.org/x/crypto/bcrypt"
)

//HashString hashes and salt a string
func HashString(pass string) (string, error) {
	var err error

	//Hash and salt string to bytes
	key, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	//Return bytes to string
	return string(key), nil
}

//MatchString compare a string to a hash
func MatchString(pass string, hash string) bool {
	var err error

	//Match string to hash
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err == nil {
		return true
	}

	return false
}
