package routes

import (
	"github.com/gorilla/mux"
	"github.com/loremcookie/go-home-pirelay-backend/internal/api/controller"
	"github.com/loremcookie/go-home-pirelay-backend/internal/api/middleware"
	"net/http"
)

//SetRoutes sets up routes and middleware
func SetRoutes() *mux.Router {
	//Create the router
	r := mux.NewRouter()
	//Register sub router that requires admin
	admin := r.PathPrefix("/api/admin/").Subrouter()

	//Here goes the middleware to use.
	//eg. r.Use(middleware)

	//AuthenticationMiddleware is responsible for regular jwt authentication
	r.Use(middleware.AuthenticationMiddleware)
	//Admin middleware is responsible for making the admin sub router only accessible for admin accounts
	admin.Use(middleware.AdminMiddleware)

	//Register functions here
	//eg. r.HandleFunc("SomePath", someHandlerFunc)
	r.HandleFunc("/api/ping", controller.PongGET).Methods(http.MethodGet)         //Only accept GET
	r.HandleFunc("/api/login", controller.LoginPOST).Methods(http.MethodPost)     //Only accept POST
	r.HandleFunc("/api/refresh", controller.RefreshPOST).Methods(http.MethodPost) //Only accept PORT

	//Garage handlers
	r.HandleFunc("/api/garage/SendSignal", controller.GarageSendSignalGET).Methods(http.MethodGet) //Only accept GET
	r.HandleFunc("/api/garage/GetState", controller.GarageGetStatusWS)                             //Only accepts web socket connections

	//Register admin routes
	admin.HandleFunc("/users/GetUser", controller.GetUserGET).Methods(http.MethodGet)       //Only accept GET request
	admin.HandleFunc("/users/GetAll", controller.GetAllUserGET).Methods(http.MethodGet)     //Only accept GET request
	admin.HandleFunc("/users/NewUser", controller.CreateUserPOST).Methods(http.MethodPost)  //Only accept POST request
	admin.HandleFunc("/users/DeleteUser", controller.DeleteUserGET).Methods(http.MethodGet) //Only accept POST request

	return r
}
