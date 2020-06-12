package main

import (
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"github.com/loremcookie/go-home/backend/internal/api/models"
	"github.com/loremcookie/go-home/backend/internal/api/routes"
	"github.com/loremcookie/go-home/backend/internal/database"
	"github.com/loremcookie/go-home/backend/internal/passhash"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	var err error

	//Load config file into environment variables to access in all parts of the program.
	//The config can be accessed just like normal environment variables with os.Getenv("NAME_OF_VARIABLE")
	err = godotenv.Load("./config/api/API_CONFIG.env")
	//Exit and log error when loading of config fails
	if err != nil {
		log.Fatalln(err)
	}

	//Convert string returned by os.Getenv() to int for generation of time.Duration
	intTimeoutBoltDB, err := strconv.Atoi(os.Getenv("BOLTDB_TIMEOUT"))
	if err != nil {
		log.Fatalln(err)
	}
	//Configure database
	err = database.Configure(&database.Database{
		Path:    os.Getenv("BOLTDB_LOCATION"),
		Timeout: time.Duration(intTimeoutBoltDB) * time.Second,
	})
	if err != nil {
		log.Fatalln(err)
	}

	//Create default user for development
	//TODO: Remove user after development
	defaultUsername := os.Getenv("DEFAULT_USER_USERNAME")
	defaultPassword, err := passhash.HashString(os.Getenv("DEFAULT_USER_PASSWORD"))
	if err != nil {
		log.Fatalln(err)
	}
	err = models.NewUser(&models.User{
		Username: defaultUsername,
		Password: defaultPassword,
		Admin:    true,
	})

	//Set up routes and middleware
	r := routes.SetRoutes()

	//Set http Handler to mux.Router
	http.Handle("/", r)

	//Start http server with config and log error
	log.Fatal(http.ListenAndServe(os.Getenv("API_HOST")+":"+os.Getenv("API_PORT") /*Format port and host to listen on*/, handlers.RecoveryHandler() /*Set panic recovery*/ (handlers.LoggingHandler(os.Stdout, r)) /*Set logging handler*/))
}
