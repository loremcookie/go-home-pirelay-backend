package jsplugin

import (
	"encoding/json"
	"fmt"
	"github.com/loremcookie/go-home/backend/internal/api/models"
	"os"
	"path/filepath"
)

//getPluginInfo retrieves the info about a plugin over decoding the plugins.json file
//that should be included in every plugin.
func getPluginInfo(pluginName string) (*models.PluginConf, error) {
	var err error

	//Create plugin info tp decode json to
	var pluginInfo models.PluginConf

	//Open plugin.json and decode info
	file, err := os.Open(filepath.FromSlash(fmt.Sprintf("%s/plugin.json", pluginName)))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//Decode json
	err = json.NewDecoder(file).Decode(&pluginInfo)
	if err != nil {
		return nil, err
	}

	//Return plugin info for further use
	return &pluginInfo, nil
}
