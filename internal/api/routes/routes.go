package routes

import (
	"github.com/gorilla/mux"
	"github.com/loremcookie/go-home/backend/internal/api/controller"
	"github.com/loremcookie/go-home/backend/internal/api/middleware"
	"net/http"
)

//SetRoutes sets up routes and middleware
func SetRoutes() {
	//Create the router
	r := mux.NewRouter()
	//Register subrouter that requires admin
	admin := r.PathPrefix("/api/admin/").Subrouter()

	//Here goes the middleware to use.
	//eg. r.Use(middleware)

	//AuthenticationMiddleware is responsible for regular jwt authentication
	r.Use(middleware.AuthenticationMiddleware)
	//Admin middleware is responsible for making the admin subrouter only accessible for admin accounts
	admin.Use(middleware.AdminMiddleware)

	//Register functions here
	//eg. r.HandleFunc("SomePath", someHandlerFunc)
	r.HandleFunc("/api/ping", controller.PongGET).Methods(http.MethodGet)         //Only accept GET
	r.HandleFunc("/api/login", controller.LoginPOST).Methods(http.MethodPost)     //Only accept POST
	r.HandleFunc("/api/refresh", controller.RefreshPOST).Methods(http.MethodPost) //Only accept PORT

	//Set http Handler to mux.Router
	http.Handle("/", r)
}
