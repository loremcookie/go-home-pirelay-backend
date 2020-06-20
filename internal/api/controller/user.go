package controller

import (
	"github.com/loremcookie/go-home/backend/internal/api/authentication"
	"github.com/loremcookie/go-home/backend/internal/api/models"
	"github.com/loremcookie/go-home/backend/internal/webutil"
	"net/http"
)

//CreateUserPOST creates a user from a user model in a request.
func CreateUserPOST(w http.ResponseWriter, r *http.Request) {
	var err error

	//Create user model to parse request body into
	var user models.User

	//Parse request body
	err = webutil.ParseReq(r, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	//Validate input
	if user.Username == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	if user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	if user.Admin == false {
		w.WriteHeader(http.StatusBadRequest)
	}

	//Create new user
	err = models.NewUser(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	//Return 200 status code when all has gone successfully
	w.WriteHeader(http.StatusOK)
}

//GetUserPOST gets a user info based os the username
func GetUserPOST(w http.ResponseWriter, r *http.Request) {
	var err error

	//Make map to decode request body to
	var reqMap map[string]interface{}

	//Decode json to map
	err = webutil.ParseReq(r, &reqMap)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Validate input
	if reqMap["username"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Check if requested user matches the user from the token
	//so you can requested your own user info but not other user
	//information
	//Get token from header
	token, err := authentication.GetTokenFromHeader(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	//Get token metadata
	tokenMetadata := authentication.GetTokenMetadata(token)
	if tokenMetadata == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	//Check if requested username matches token username if user is not admin
	if tokenMetadata.Admin == false && tokenMetadata.Username != reqMap["username"] {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	//Get user by username
	user, err := models.GetUser(reqMap["username"].(string))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Respond to request with user info
	err = webutil.Respond(w, 200, map[string]interface{}{
		"user": user,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Return 200 status code when all has gone successfully
	w.WriteHeader(http.StatusOK)
}

//DeleteUserPOST deletes a user by the username
func DeleteUserPOST(w http.ResponseWriter, r *http.Request) {
	var err error

	//Make map to parse request into
	var reqMap map[string]interface{}

	//Parse request json into map
	err = webutil.ParseReq(r, reqMap)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Validate input
	if reqMap["username"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Delete user
	err = models.DeleteUser(reqMap["username"].(string))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Return 200 status code when all has gone successfully
	w.WriteHeader(http.StatusOK)
}

//GetAllUserGET retrieves all user registers in the USER bucket
func GetAllUserGET(w http.ResponseWriter, r *http.Request) {
	var err error

	//Get all users as slice
	users := models.GetAllUser()

	//Respond to request with users
	err = webutil.Respond(w, 200, map[string]interface{}{
		"users": users,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	//Return 200 status code when all has gone successfully
	w.WriteHeader(http.StatusOK)
}
