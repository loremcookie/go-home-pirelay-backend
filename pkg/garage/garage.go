package garage

import (
	"github.com/loremcookie/go-home/backend/pkg/gpio"
	"github.com/loremcookie/go-home/backend/pkg/pirelay"
	"os"
	"strconv"
	"time"
)

//SendSignal sends a signal to the garage to open, close or hold depending on the current state
func SendSignal() error {
	var err error

	//Get relay used for garage and convert it to int
	garageRelay, err := strconv.Atoi(os.Getenv("Garage_Relay"))
	if err != nil {
		return err
	}

	//Open the garage relay
	err = pirelay.OpenRelay(garageRelay)
	if err != nil {
		return err
	}

	//Sleep 600 Milliseconds for garage to realise the relay has been opened
	time.Sleep(600 * time.Millisecond)

	//Close relay again
	err = pirelay.CloseRelay(garageRelay)
	if err != nil {
		return err
	}

	return nil
}

//GetStatus returns the current status of the garage open or not open
func GetStatus() (string, error) {
	var err error

	//Get environment variable and convert it from string to int
	sensorPin, err := strconv.Atoi(os.Getenv("Garage_Pin_Sensor"))
	if err != nil {
		return "", err
	}

	//Get state of sensor pin
	state, err := gpio.GetState(sensorPin)
	if err != nil {
		return "", err
	}

	//Switch on the pin state to match garage state to
	switch state {
	case gpio.High:
		//If current is on the pin the garage is closed because the circuit
		//is closed
		return "closed", nil
	case gpio.Low:
		//If no current is on the pin the garage is open because the circuit
		//is broken
		return "open", nil
	}

	return "", nil
}
