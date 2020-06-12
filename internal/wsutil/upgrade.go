package wsutil

import (
	"github.com/gorilla/websocket"
	"net/http"
)

//WsHandler is the type of function that is accepted by the UpgradeWS functions for handling the websocket connection
type WsHandler func(conn *websocket.Conn) error

//upgrader is the configuration for the buffer sizes of the websocket connection
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//UpgradeWS upgrades a standard http endpoint to a websocket connection
func UpgradeWS(w http.ResponseWriter, r *http.Request, wsHandler WsHandler) error {
	var err error

	//Upgrade connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	defer conn.Close() //Close connection when function ends

	//Call websocket handler and return error when returned from websocket handler
	err = wsHandler(conn)
	if err != nil {
		return err
	}

	return nil
}
