package app

import (
	"github.com/fahribaharudin/api_gateway/app/controllers"
)

// BaseController ..
type BaseController struct {
	APIGatewayController controllers.APIGatewayController
	HomeController       controllers.HomeController
}

var baseController BaseController

func init() {
	baseController = BaseController{
		APIGatewayController: controllers.APIGatewayController{},
		HomeController:       controllers.HomeController{},
	}
}

// RegisterRoutes register non api gateway routes
func (app *Kernel) RegisterRoutes() {
	router := app.APIGatewayRouter.Router

	// handling non api gateway router
	router.HandleFunc("/", baseController.HomeController.LandingEndpointHandler).Methods("GET")
	// router.GET("/", baseController.HomeController.LandingEndpointHandler)

	app.APIGatewayRouter.Router = router
}
