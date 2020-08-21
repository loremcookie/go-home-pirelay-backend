package gpio

import "testing"

//TestSetState is the test function for SetState
func TestSetState(t *testing.T) {
	var err error

	//pinNum is the number of the pin that is tested
	pinNum := 1

	//Set state of pin to on
	err = SetState(pinNum, int(High))
	if err != nil {
		t.Error(err)
	}

	//Get state of pin that should be high
	pinState, err := GetState(pinNum)
	if err != nil {
		t.Error(err)
	}

	//Check if state of pin is set
	if pinState != High {
		t.Error("Pin status failed to set")
	}
}

//TestGetState is the unit test for the GetState function
func TestGetState(t *testing.T) {
	var err error

	//pinNum is the number of the pin that is tested
	pinNum := 1

	//Set state of pin to on
	err = SetState(pinNum, int(High))
	if err != nil {
		t.Error(err)
	}

	//Get state of pin that should be high
	pinState, err := GetState(pinNum)
	if err != nil {
		t.Error(err)
	}

	//Check if state of pin is set
	if pinState != High {
		t.Error("Pin status failed to set")
	}
}
