package controller

import (
	"github.com/loremcookie/go-home/backend/internal/api/webutil"
	"net/http"
)

//PongGET is a test route to check if the api is up.
//If you send a request to this route then you get a pong.
func PongGET(w http.ResponseWriter, r *http.Request) {
	var err error

	//Declare response
	resp := map[string]string{
		"ping": "pong",
	}

	err = webutil.Respond(w, http.StatusOK, resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
