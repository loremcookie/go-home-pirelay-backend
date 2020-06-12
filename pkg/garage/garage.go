/*
TODO: uncomment actual logic after development
This package is a layer for the relay package responsible for the garage.
With this package you can send a signal to the garage and receive the current status of the garage eg. Open and Not Open.
*/
package garage

import "fmt"

//SendSignal sends a signal to the garage to open, close or hold depending on the current state
func SendSignal() error {
	//var err error

	////Get relay used for garage and convert it to int
	//garageRelay, err := strconv.Atoi(os.Getenv("Garage_Relay"))
	//if err != nil {
	//	return err
	//}

	////Open the garage relay
	//err = relay.OpenRelay(garageRelay)
	//if err != nil {
	//	return err
	//}

	////Sleep 600 Milliseconds for garage to realise the relay has been opened
	//time.Sleep(600 * time.Millisecond)

	////Close relay again
	//err = relay.CloseRelay(garageRelay)
	//if err != nil {
	//	return err
	//}

	//return nil

	fmt.Println("Send Signal to garage")
	return nil
}

//GetStatus returns the current status of the garage open or not open
func GetStatus() (string, error) {
	//var err error

	////Open gpio memory range
	//err = rpio.Open()
	//if err != nil {
	//	return "", err
	//}
	//defer rpio.Close() //Close gpio memory range when function ends

	////Get environment variable and convert it from string to int
	//sensorPin, err := strconv.Atoi(os.Getenv("Garage_Pin_Sensor"))
	//if err != nil {
	//	return "", err
	//}

	////Read from pin
	//pin := rpio.Pin(sensorPin)
	//pin.Input()
	//state := pin.Read()
	//switch state {
	//case rpio.High:
	//	return StateClosed, nil
	//case rpio.Low:
	//	return "open", nil
	//}
	//return "", nil

	return "closed", nil
}
