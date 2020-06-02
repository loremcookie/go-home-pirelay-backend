package models

//Login model is the model used to log in.
//Its contains the information used to login is more info is added to log in
//you need to add it here.
type Login struct {
	Username string
	Password string
}

//TokenClaims are the private claims stored in the token.
//If private claims are added you need to add them here.
type TokenClaims struct {
	Username string
	Admin    bool
}

//PluginConf is the json struct to decode the plugin info from
//the plugin.json from the plugin directory.
type PluginConf struct {
	Name     string          `json:"name"`     //Name of the main plugin script file should be the name of the plugin
	Author   string          `json:"author"`   //Author of the plugin
	Version  string          `json:"version"`  //Is the version of the plugin
	Handlers []PluginHandler `json:"handlers"` //This is a list of handlers to register handler
	EnvVars  []string        `json:"env-vars"` //These are the environment variables the plugin can access
}

//PluginHandler is the specification to register a handler with a plugin
type PluginHandler struct {
	HandlerFunc string `json:"handler-func"` //The function that handles the request
	Route       string `json:"route"`        //Is the route the functions handles
	Method      string `json:"method"`       //Method that is allowed
}
