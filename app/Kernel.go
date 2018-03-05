package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

/************************************/
/********** ROUTER ENGINE ***********/
/************************************/

// RoutesWrapper is the main object of application router
type RoutesWrapper struct {
	Endpoints []Endpoint
	Router    *mux.Router
}

// Endpoint object to model some specific endpoints
type Endpoint struct {
	Path    string
	Method  string
	Handler interface{}
}

// AddEndpoint used to define new endpoint to the app
func (r *RoutesWrapper) AddEndpoint(path string, method string, handler ...interface{}) {
	var endpoint = Endpoint{Path: path, Method: method}
	if len(handler) > 0 {
		endpoint.Handler = handler[0]
	}

	r.Endpoints = append(r.Endpoints, endpoint)
}

// Handle the dispatching process
func (r *RoutesWrapper) Handle() {
	// iterate over the stored endpoints
	for _, route := range r.Endpoints {
		r.Router.HandleFunc(route.Path, baseController.APIGatewayController.UniversalHandler).Methods(
			"GET", "POST", "PUT", "PATCH", "DELETE",
		)
	}
}

/************************************/
/******* KERNEL / APPLICATION *******/
/************************************/

// Kernel is the main app wrapper object
type Kernel struct {
	APIGatewayRouter RoutesWrapper
}

// Init is the Kernel constructor ..
func (app *Kernel) Init() {
	app.APIGatewayRouter = RoutesWrapper{}
	app.APIGatewayRouter.Router = mux.NewRouter()
}

// ParseSwaggerAPIEndpoints is parsing swagger file and register it routes to the api gateway
func (app *Kernel) ParseSwaggerAPIEndpoints() {

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

			app.APIGatewayRouter.AddEndpoint(path, method)
		}
	}
}

// Run the app!
func (app *Kernel) Run() {

	// handling api gateway router
	app.APIGatewayRouter.Handle()

	// add some middleware
	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		// before middleware..
		next(rw, r)
		// after middleware..
	}))
	n.UseHandler(app.APIGatewayRouter.Router)

	// run the app
	server := http.Server{
		Addr:         ":8000",
		Handler:      n,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Println("Server running on http://localhost:8000")
	server.ListenAndServe()
}
