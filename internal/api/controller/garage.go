package controller

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/loremcookie/go-home-pirelay-backend/internal/wsutil"
	"github.com/loremcookie/go-home-pirelay-backend/pkg/garage"
	"net/http"
	"time"
)

//GarageSendSignalGET sends a signal to the garage to open, close or stop the garage motor
func GarageSendSignalGET(w http.ResponseWriter, _ *http.Request) {
	var err error

	//Send signal to garage motor
	err = garage.SendSignal()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	//When everything goes right send 200 status code
	w.WriteHeader(http.StatusOK)
}

//GarageGetStatusWS is the websocket endpoint that is responsible for serving the
///status of the garage in realtime (1 second delay by updated status)
func GarageGetStatusWS(w http.ResponseWriter, r *http.Request) {
	var err error

	//Upgrade connection to web socket connection
	err = wsutil.UpgradeWS(w, r, func(conn *websocket.Conn) error {
		//Send every second a status update in the garage
		for {
			//Get status of garage
			status, err := garage.GetStatus()
			if err != nil {
				return err
			}

			//Send status of garage to client
			err = wsutil.RespondWS(conn, map[string]interface{}{
				"state": status,
			})
			if err != nil {
				return err
			}

			//Sleep 1 second to dont stress the cpu as much
			time.Sleep(1 * time.Second)
		}
	})
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

}
