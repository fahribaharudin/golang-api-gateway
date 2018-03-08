package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/negroni"
)

// RegisterRoutes to handle the given endpoints
func (app *Kernel) RegisterRoutes() {
	app.Router.HandleFunc("/", app.Controllers.APIGateway.WelcomeHandler)

	// some router with middleware
	app.Router.PathPrefix("/admin").Handler(negroni.New(negroni.HandlerFunc(app.Middlewares().AdminMiddleware)))

	// registering the petstore endpoint from swagger api
	var swaggerAPI map[string]interface{}
	jsonFileContent, err := ioutil.ReadFile("./petstore.swagger.json")
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}

	err = json.Unmarshal(jsonFileContent, &swaggerAPI)
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}

	for path, pathDefinitions := range swaggerAPI["paths"].(map[string]interface{}) {
		for method, definition := range pathDefinitions.(map[string]interface{}) {
			deprecated := definition.(map[string]interface{})["deprecated"]
			if deprecated != nil && deprecated == true {
				continue
			}

			app.Router.HandleFunc(path, app.Controllers.APIGateway.UniversalHandler).Methods(method)
		}
	}
}
