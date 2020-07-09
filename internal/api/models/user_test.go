package models

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/loremcookie/go-home-pirelay-backend/internal/database"
	"os"
	"testing"
	"time"
)

var (
	//Test data for user
	testUser = User{
		Username: "test_user",
		Password: "test_pass",
		Admin:    false,
	}
)

//Main function of the test prepare and tear down the database
func TestMain(m *testing.M) {
	var err error

	//Prepare tests and initialize database
	//Check is TESTDATA_DIR directory exists
	//The TESTDATA_DIR directory is used to hold data used in tests
	if _, err = os.Stat(os.Getenv("TESTDATA_DIR")); err != nil {
		//When TESTDATA_DIR does not exist then create it
		if err == os.ErrNotExist {
			err = os.Chdir(os.Getenv("TESTDATA_DIR"))
			//If creation of directory fails fail test
			if err != nil {
				fmt.Printf("Test directory %s could not be created.", os.Getenv(""))
				os.Exit(2)
			}
		}
	}

	//Init database package to be used in the model test
	err = database.Configure(&database.Database{
		Path:    fmt.Sprintf("%s/TEST_API_DB.db", os.Getenv("TESTDATA_DIR")),
		Timeout: 5 * time.Second,
	})
	//If database configuration fails fail test
	if err != nil {
		fmt.Println("Unable to initialize database.")
		os.Exit(2)
	}

	//Run tests and retrieve exit value
	exitVal := m.Run()

	//Clean up after tests
	//Close database connection
	err = database.Close()
	if err != nil {
		fmt.Println("Unable to close database connection.")
		os.Exit(2)
	}

	//Remove TESTDATA_DIR directory if unable to remove the directory exit
	err = os.Remove(os.Getenv("TESTDATA_DIR"))
	if err != nil {
		fmt.Printf("Unable to remove %s directory.", os.Getenv("TESTDATA_DIR"))
		os.Exit(2)
	}

	//Exit with exit values of test
	os.Exit(exitVal)
}

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

	//Reset database
	err = database.Reset()
	if err != nil {
		t.Error(err)
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

	//Reset database
	err = database.Reset()
	if err != nil {
		t.Error(err)
	}
}

//Test DeleteUser is the test function for DeleteUser
func TestDeleteUser(t *testing.T) {
	var err error

	//Create user to delete
	err = NewUser(&testUser)
	if err != nil {
		t.Error(err)
	}

	//Delete user
	err = DeleteUser(testUser.Username)
	if err != nil {
		t.Error(err)
	}

	//Validate that user is deleted
	user, err := GetUser(testUser.Username)
	//Check that user is deleted
	if err != nil {
		if err != bolt.ErrInvalid {
			t.Error(err)
		}
	}
	if user != nil {
		t.Error("User must be nil")
	}
}

//TestValidUser test for the GetUser function
func TestValidUser(t *testing.T) {
	var err error

	//Make test user
	err = NewUser(&testUser)
	if err != nil {
		t.Error(err)
	}

	//Validate user with data that is valid
	isValid := ValidUser(&Login{
		Username: testUser.Username,
		Password: testUser.Password,
	})

	//User must be valid
	if !isValid {
		t.Error("Valid user data is marked as not valid")
	}

	//Reset database
	err = database.Reset()
	if err != nil {
		t.Error("Unable to reset database")
	}
}
