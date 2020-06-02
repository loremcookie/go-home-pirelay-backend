package jsplugin

import (
	"bytes"
	"github.com/robertkrimen/otto"
	"os"
)

var (
	runtime *otto.Otto //Runtime is the javascript vm runtime
)

//newRuntime makes a new runtime sets it as global variable and loads a file into it.
func newRuntime(filename string) error {
	var err error

	//Open script to load into runtime
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	//New buffer to read file into
	var buf bytes.Buffer

	//Read file into buffer
	_, err = buf.ReadFrom(file)
	if err != nil {
		return err
	}

	//Make new runtime
	runtime = otto.New()

	//Load buffer into runtime
	_, err = runtime.Run(buf.String())
	if err != nil {
		return err
	}

	return nil
}
