package webutil

import (
	"encoding/json"
	"net/http"
)

//ParseReq parses json out of the request body
func ParseReq(r *http.Request, v interface{}) error {
	var err error

	//Decode json request body
	err = json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		return err
	}
	defer r.Body.Close() //Close request body when function ends

	return nil
}
