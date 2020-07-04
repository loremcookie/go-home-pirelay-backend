package models

import "testing"

var (
	//Test data for user
	testUser = User{
		Username: "test_user",
		Password: "test_pass",
		Admin:    false,
	}
)

//TestNewUser test NewUser function
func TestNewUser(t *testing.T) {
	var err error

	//Create new user
	err = NewUser(&testUser)
	if err != nil {
		t.Error(err)
	}

	//Validate that user was created
	userValid, err := GetUser(testUser.Username)
	if err != nil {
		t.Error(err)
	}

	if testUser != *userValid {
		t.Error("Created user does not match database user: \nTestData:", testUser, "\nDatabaseUser:", userValid)
	}
}

//TestGetUser test for the GetUser function
func TestGetUser(t *testing.T) {
	var err error

	//Create new user
	err = NewUser(&testUser)
	if err != nil {
		t.Error(err)
	}

	//Validate creation of user
	userValid, err := GetUser(testUser.Username)
	if err != nil {
		t.Error(err)
	}

	if testUser != *userValid {
		t.Error("Created user does not match database user: \nTestData:", testUser, "\nDatabaseUser:", userValid)
	}
}
