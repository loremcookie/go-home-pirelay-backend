package pirelay

import (
	"github.com/stianeikeland/go-rpio"
	"os"
	"strconv"
)

//GetPinByRelay returns the pin by the relay
func GetPinByRelay(relay int) (int, error) {
	var err error
	matchRelayPin := make(map[int]int)
	matchRelayPin[1], err = strconv.Atoi(os.Getenv("1_RELAY_PIN"))
	if err != nil {
		return 0, err
	}
	matchRelayPin[2], err = strconv.Atoi(os.Getenv("2_RELAY_PIN"))
	if err != nil {
		return 0, err
	}
	matchRelayPin[3], err = strconv.Atoi(os.Getenv("3_RELAY_PIN"))
	if err != nil {
		return 0, err
	}
	matchRelayPin[4], err = strconv.Atoi(os.Getenv("4_RELAY_PIN"))
	if err != nil {
		return 0, err
	}
	return matchRelayPin[relay], nil
}

//OpenRelay Opens the specified relay
//Doesnt work when relay is already open
func OpenRelay(relay int) error {
	var err error

	//Open gpio memory range
	err = rpio.Open()
	if err != nil {
		return err
	}
	defer rpio.Close() //Close gpio memory range when function ends

	//Get Pin by relay
	relayPin, err := GetPinByRelay(relay)
	if err != nil {
		return err
	}

	//Output to the relay pin to open the relay up
	pin := rpio.Pin(relayPin)
	pin.Mode(rpio.Output)
	pin.Write(rpio.High)

	return nil
}

//CloseRelay Closed the specified relay
//Doesnt Work when relay is already closed
func CloseRelay(relay int) error {
	var err error

	//Open gpio memory range
	err = rpio.Open()
	if err != nil {
		return err
	}
	defer rpio.Close() //Close gpio memory range when function ends

	//Get Pin by relay
	relayPin, err := GetPinByRelay(relay)
	if err != nil {
		return err
	}

	//Stop Output to pin to close relay
	pin := rpio.Pin(relayPin)
	pin.Write(rpio.Low)

	return nil
}
