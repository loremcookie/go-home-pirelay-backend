package jsplugin

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/loremcookie/go-home/backend/internal/api/models"
	"github.com/robertkrimen/otto"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

//LoadPlugins recursively loads plugins in a specific folder
func LoadPlugins(pluginFolder string, router *mux.Router) error {
	var err error

	//Read directory and get all files and folders
	files, err := ioutil.ReadDir(pluginFolder)
	if err != nil {
		return err
	}

	//Create lis to store all plugin directory's in
	var pluginDirs []string

	//Loop through files
	for _, file := range files {

		//Check if file is directory
		if file.IsDir() {
			//When file is directory append it to list of directory's
			dirPath := fmt.Sprintf("%s/%s", pluginFolder, file.Name())
			pluginDirs = append(pluginDirs, dirPath)
		}
	}

	//Range through directory's and load them with recursion
	for _, dir := range pluginDirs {

		//Load current plugin
		err = LoadPlugin(dir, router)
		if err != nil {
			log.Printf("Err loading plugin %s\n", dir)
		}

		//Load next plugins in subdirectory's
		err = LoadPlugins(dir, router)
		if err != nil {
			log.Printf("Err loading plugin %s\n", dir)
		}
	}

	return nil
}

//LoadPlugin loads a plugin and registers api functions
func LoadPlugin(pluginName string, router *mux.Router) error {
	var err error

	//Get plugin info
	pluginInfo, err := getPluginInfo(pluginName)
	if err != nil {
		return err
	}

	//Make a new runtime
	err = newRuntime(filepath.FromSlash(fmt.Sprintf("%s/%s.js", pluginName, pluginInfo.Name)))
	if err != nil {
		return err
	}

	//Check if plugin is correctly loaded
	if runtime == nil {
		log.Printf("Error lading plugin %s \n", pluginInfo.Name)
		return nil
	}

	//Load api functions into runtime
	err = loadApiFunc(pluginInfo)
	if err != nil {
		return err
	}

	//Register plugin handler
	registerPluginHandler(pluginInfo, router)

	return nil
}

//loadApiFunc loads api functions into javascript runtime
func loadApiFunc(pluginInfo *models.PluginConf) error {
	var err error

	//Load environment variables into js runtime
	err = loadEnvVariables(pluginInfo)
	if err != nil {
		return err
	}

	//Register api functions to interact with the program
	err = runtime.Set("log", logJS) //Log a message to stdout
	if err != nil {
		return err
	}

	err = runtime.Set("ParseReq", parseReqJS) //Parse json request body
	if err != nil {
		return err
	}

	err = runtime.Set("SendResp", sendRespJS) //Send json response
	if err != nil {
		return err
	}

	err = runtime.Set("OpenRelay", openRelayJS) //Open relay
	if err != nil {
		return err
	}

	err = runtime.Set("CloseRelay", closeRelayJS) //Close relay
	if err != nil {
		return err
	}

	return nil
}

//loadEnvVariables loads and registers environment variables to javascript runtime
func loadEnvVariables(pluginInfo *models.PluginConf) error {
	var err error

	//Make map to store environment key value pairs in
	var envs = map[string]interface{}{}

	//Range through envKeys and assign values to keys
	for _, envKey := range pluginInfo.EnvVars {
		envVal := os.Getenv(envKey) //Get environment value from key
		envs[envKey] = envVal       //Assign key and value to map
	}

	//Make function to get environment variables
	//It is only for use in js vm
	envJS := func(key string) otto.Value {
		envVal, ok := envs[key]
		if !ok {
			return otto.Value{}
		}

		//Convert string to otto value
		retVar, err := runtime.ToValue(envVal)
		if err != nil {
			return otto.Value{}
		}
		return retVar
	}

	//Register env function
	err = runtime.Set("env", envJS)
	if err != nil {
		return err
	}

	return nil
}

//registerPluginHandler registers the http handler associated with plugins
func registerPluginHandler(pluginInfo *models.PluginConf, router *mux.Router) {
	//Loop through plugin handlers to register them
	for _, handler := range pluginInfo.Handlers {

		//Register handler
		router.HandleFunc(handler.Route, func(w http.ResponseWriter, r *http.Request) {
			var err error

			//Convert response writer to otto value
			wJS, err := runtime.ToValue(w)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			}

			//Convert request to otto value
			rJS, err := runtime.ToValue(*r)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			}

			//Call handler js functions with appropriate arguments
			_, err = runtime.Call(handler.HandlerFunc, nil, []otto.Value{wJS, rJS})
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			}

		}).Methods(handler.Method) //Make function available through only whitelisted http methods
	}
}
