package wsutil

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/websocket"
)

//RespondWS sends a json response to a web socket connection
func RespondWS(conn *websocket.Conn, resp interface{}) error {
	var err error

	//Make buffer to encode json to
	var buf bytes.Buffer

	//Encode response to json
	err = json.NewEncoder(&buf).Encode(&resp)
	if err != nil {
		return err
	}

	//Send buffer to connection
	err = conn.WriteMessage(websocket.TextMessage, buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}
