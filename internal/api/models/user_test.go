package models

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/joho/godotenv"
	"github.com/loremcookie/go-home-pirelay-backend/internal/database"
	"github.com/loremcookie/go-home-pirelay-backend/internal/passhash"
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

	//Load test configuration file
	err = godotenv.Load("../../../config/api/TEST_CONFIG.env")
	if err != nil {
		fmt.Printf("Unable to load test config: %s", err)
		os.Exit(2)
	}

	//Prepare tests and initialize database
	//Check is TESTDATA_DIR directory exists
	//The TESTDATA_DIR directory is used to hold data used in tests
	//When TESTDATA_DIR does not exist then create it
	if _, err = os.Stat(fmt.Sprintf("../../../%s", os.Getenv("TESTDATA_DIR"))); os.IsNotExist(err) {
		err = os.Mkdir(fmt.Sprintf("../../../%s", os.Getenv("TESTDATA_DIR")), 0755)
		if err != nil {
			fmt.Printf("Test directory %s could not be created.", os.Getenv("TESTDATA_DIR"))
			os.Exit(2)
		}
	}

	//Init database package to be used in the model test
	err = database.Configure(&database.Database{
		Path:    fmt.Sprintf("../../../%s/TEST_API_DB.db", os.Getenv("TESTDATA_DIR")),
		Timeout: 5 * time.Second,
	})
	//If database configuration fails fail test
	if err != nil {
		fmt.Printf("Unable to initialize database: %s", err)
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
	err = os.RemoveAll(fmt.Sprintf("../../../%s", os.Getenv("TESTDATA_DIR")))
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

	//Reset database
	err = database.Reset()
	if err != nil {
		t.Error(err)
	}
}

//TestValidUser test for the GetUser function
func TestValidUser(t *testing.T) {
	var err error

	//Store password to be reset after test
	orgTestPass := testUser.Password

	//Hash password to be stored in the database
	testUser.Password, err = passhash.HashString(testUser.Password)
	if err != nil {
		t.Error(err)
	}

	//Make test user
	err = NewUser(&testUser)
	if err != nil {
		t.Error(err)
	}

	//Validate user with data that is valid
	isValid := ValidUser(&Login{
		Username: testUser.Username,
		Password: orgTestPass,
	})

	//User must be valid
	if !isValid {
		t.Error("Valid user data is marked as not valid")
	}

	//Reset Password
	testUser.Password = orgTestPass

	//Reset database
	err = database.Reset()
	if err != nil {
		t.Error("Unable to reset database")
	}
}
