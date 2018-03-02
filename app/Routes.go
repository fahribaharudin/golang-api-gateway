package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterNonAPIGatewayRoutes ..
func (app *Kernel) RegisterNonAPIGatewayRoutes() {
	router := app.APIGatewayRouter.Router

	// handling non api gateway router
	router.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "Welcome bro!") })

	app.APIGatewayRouter.Router = router
}
