package controller

import (
	"fmt"
	"github.com/loremcookie/go-home/backend/internal/api/webutil"
	"net/http"
)

//RefreshPOST requires a refresh token to return
//a new access token and refresh token.
func RefreshPOST(w http.ResponseWriter, r *http.Request) {
	var err error

	//Make map to store
	var reqMap map[string]string

	//Get refresh token from request
	err = webutil.ParseReq(r, &reqMap)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Validate input
	if reqMap["refresh_token"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(reqMap["refresh_token"])
	w.WriteHeader(http.StatusOK)
	return
}
