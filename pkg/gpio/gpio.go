package gpio

import (
	"errors"
	"fmt"
	"github.com/stianeikeland/go-rpio"
)

//Set types
type State uint8

//Declare the variable that are the states of the gpio pins
var (
	High State = 1 //High is the state of the gpio pin if current is on the pin
	Low  State = 0 //Low is the state of the pin if there is no current present
)

var (
	ErrNoState = errors.New("value not a state")
)

//SetState sets the state of a gpio pin.
//it accepts an int value 1 for High also current 0 also no current.
func SetState(pinNum int, state int) error {
	var err error

	//Validate function input
	if state != 1 || state != 0 {
		return ErrNoState
	}

	//Open gpio memory range
	err = rpio.Open()
	if err != nil {
		return err
	}
	defer rpio.Close() //Close gpio memory range when function ends

	//Convert pin Number to pin
	pin := rpio.Pin(pinNum)
	//Set pin mode to write
	pin.Mode(rpio.Output)
	//Write mode to gpio pin
	pin.Write(rpio.State(state))

	fmt.Println(pinNum)
	return nil
}

//GetState returns the state of a pio pin
//possible return values are High=1 or Low=0 it depends of the state og the pin
//if the pin has current on it High is returned when no current is present Low is returned.
func GetState(pinNum int) (State, error) {
	var err error

	//Open gpio memory range
	err = rpio.Open()
	if err != nil {
		return Low, err
	}
	defer rpio.Close() //Close gpio memory range when funtion ends

	//Convert pin number to pin
	pin := rpio.Pin(pinNum)
	//Take input from pin
	pin.Input()
	//Read from pin to get state
	state := pin.Read()

	//Assign state depending of the value of the Read return
	switch state {
	case rpio.High:
		return High, nil
	case rpio.Low:
		return Low, nil
	}

	return Low, nil
}
