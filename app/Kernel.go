package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/fahribaharudin/api_gateway/app/controllers"
	"github.com/gin-gonic/gin"
)

/************************************/
/******* KERNEL / APPLICATION *******/
/************************************/

// Kernel is the main app wrapper object
type Kernel struct {
	APIGatewayRouter RoutesWrapper
}

// Construct is the Kernel constructor ..
func (app *Kernel) Construct() {
	app.APIGatewayRouter = RoutesWrapper{}
	app.APIGatewayRouter.Router = gin.Default()
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
		for method := range pathDefinitions.(map[string]interface{}) {
			app.APIGatewayRouter.AddEndpoint(path, method)
		}
	}
}

// Run the app!
func (app *Kernel) Run() {

	// add some middleware
	app.APIGatewayRouter.Router.Use(func() gin.HandlerFunc {
		return func(c *gin.Context) {
			log.Println("Middleware start..")

			c.Next()

			log.Println("MIddleware end..")
		}
	}())

	// handling api gateway router
	app.APIGatewayRouter.Handle()

	// run the app
	app.APIGatewayRouter.Router.Run(":8000")
}

/************************************/
/********** ROUTER ENGINE ***********/
/************************************/

// RoutesWrapper is the main object of application router
type RoutesWrapper struct {
	Endpoints []Endpoint
	Router    *gin.Engine
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
		APIGatewayController := controllers.APIGatewayController{}
		switch strings.ToLower(route.Method) {
		case "get":
			r.Router.GET(route.Path, APIGatewayController.GetHandler)
			break

		case "post":
			r.Router.POST(route.Path, APIGatewayController.PostHandler)
			break

		case "put":
			r.Router.PUT(route.Path, APIGatewayController.PostHandler)
			break

		case "patch":
			r.Router.PATCH(route.Path, APIGatewayController.PatchHandler)

		case "delete":
			r.Router.DELETE(route.Path, APIGatewayController.DeleteHandler)
			break

		default:
			log.Println("The http method: " + route.Method + " is not supported by this api gateway")
			break
		}
	}
}
