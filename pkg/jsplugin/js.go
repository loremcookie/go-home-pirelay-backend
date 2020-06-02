package jsplugin

import (
	"github.com/loremcookie/go-home/backend/internal/api/webutil"
	"github.com/loremcookie/go-home/backend/pkg/pirelay"
	"github.com/robertkrimen/otto"
	"log"
	"net/http"
)

//In this file go all functions that are used in the javascript vm.
//All functions that end with the letters JS are wrappers for golang functions to function in javascript

//logJS logs a value to stdout
func logJS(statement interface{}) {
	log.Println(statement)
}

//openRelayJS opens a relay
func openRelayJS(relay int) otto.Value {
	var err error

	//Initialize return value
	var retVar otto.Value

	//Open relay
	err = pirelay.OpenRelay(relay)
	if err != nil {
		//Convert golang var to js var
		retVar, err = runtime.ToValue(false)
		if err != nil {
			log.Println(ErrTypeConversion)
			return otto.Value{}
		}
		return retVar //Return js var
	}

	//Convert golang var to js var
	retVar, err = runtime.ToValue(true)
	if err != nil {
		log.Println(ErrTypeConversion)
		return otto.Value{}
	}
	return retVar //Return js var
}

//closeRelayJS closes a relay
func closeRelayJS(relay int) otto.Value {
	var err error

	//Initialize return value
	var retVar otto.Value

	//Close relay
	err = pirelay.CloseRelay(relay)
	if err != nil {
		//Convert golang var to js var
		retVar, err = runtime.ToValue(false)
		if err != nil {
			log.Println(ErrTypeConversion)
			return otto.Value{}
		}
		return retVar
	}

	//Convert golang var to js var
	retVar, err = otto.ToValue(true)
	if err != nil {
		//Convert golang var to js var
		log.Println(ErrTypeConversion)
		return otto.Value{}
	}
	return retVar
}

//sendRespJS can write a json response to response writer
func sendRespJS(w http.ResponseWriter, status int, data map[string]interface{}) otto.Value {
	var err error

	//Initialize return variable
	var retVar otto.Value

	//Marshal and send response
	err = webutil.Respond(w, status, data)
	if err != nil {
		retVar, err = runtime.ToValue(false)
		if err != nil {
			//Convert golang var to js var
			log.Println(ErrTypeConversion)
			return otto.Value{}
		}
		return retVar
	}

	//Convert golang var to js var
	retVar, err = runtime.ToValue(true)
	if err != nil {
		//Convert golang var to js var
		log.Println(ErrTypeConversion)
		return otto.Value{}
	}
	return retVar
}

//parseReqJS can parse a json request body
func parseReqJS(r *http.Request) (otto.Value, otto.Value) {
	var err error

	//Declare status return variable
	var ok = otto.Value{}

	//Check if request is nil
	if r.Body == http.NoBody {
		//Convert golang var to js var
		ok, err = runtime.ToValue(false)
		if err != nil {
			//Convert golang var to js var
			log.Println(ErrTypeConversion)
			return otto.Value{}, otto.Value{}
		}
		return otto.Value{}, ok
	}

	//Make map to parse request info to
	var dataMap map[string]interface{}

	//Parse json to map
	err = webutil.ParseReq(r, &dataMap)
	if err != nil {
		//Convert golang var to js var
		ok, err = runtime.ToValue(false)
		if err != nil {
			//Convert golang var to js var
			log.Println(ErrTypeConversion)
			return otto.Value{}, otto.Value{}
		}
		return otto.Value{}, ok
	}

	//Convert data map to otto.Value
	retMap, err := runtime.ToValue(dataMap)
	if err != nil {
		//Convert golang var to js var
		ok, err = runtime.ToValue(false)
		if err != nil {
			//Convert golang var to js var
			log.Println(ErrTypeConversion)
			return otto.Value{}, otto.Value{}
		}
		return otto.Value{}, ok
	}
	//Convert golang var to js var
	ok, err = runtime.ToValue(true)
	if err != nil {
		//Convert golang var to js var
		log.Println(ErrTypeConversion)
		return otto.Value{}, otto.Value{}
	}
	return retMap, ok
}
