package webutil

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

//Respond sends json data to api client
func Respond(w http.ResponseWriter, status int, resp interface{}) error {
	var err error

	//Make new buffer to write json to
	var buf bytes.Buffer

	//Decode json to buffer
	err = json.NewEncoder(&buf).Encode(&resp)
	if err != nil {
		return err
	}

	//Set header to application/json
	w.Header().Set("Content-Type", "application/json")

	//Write status code
	w.WriteHeader(status)

	//Copy buffer to ResponseWrite
	_, err = io.Copy(w, &buf)
	if err != nil {
		return err
	}

	return nil
}
