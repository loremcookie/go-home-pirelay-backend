package controller

import (
	"github.com/loremcookie/go-home/backend/internal/api/authentication"
	"github.com/loremcookie/go-home/backend/internal/api/models"
	"github.com/loremcookie/go-home/backend/internal/api/webutil"
	"net/http"
)

//LoginPOST requires a username and password.
//It returns the jwt access token.
func LoginPOST(w http.ResponseWriter, r *http.Request) {
	var err error

	//Make empty login model to store request body
	var reqUser models.Login

	//Parse request in to empty user struct
	err = webutil.ParseReq(r, &reqUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Verify user
	if !models.ValidUser(&reqUser) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//Get user struct from username
	user, err := models.GetUser(reqUser.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Fill in token claims with user
	privateClaims := models.TokenClaims{
		Username: user.Username,
		Admin:    user.Admin,
	}

	//Create token
	tokens, err := authentication.NewToken(&privateClaims)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Create response in map
	resp := map[string]string{
		"token_type":    "bearer",
		"access_token":  tokens["access_token"],
		"refresh_token": tokens["refresh_token"],
	}

	err = webutil.Respond(w, http.StatusOK, resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
