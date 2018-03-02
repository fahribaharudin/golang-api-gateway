package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeController handler to some endpoint
type HomeController struct {
}

// LandingEndpointHandler to GET:/
func (controller *HomeController) LandingEndpointHandler(c *gin.Context) {
	c.String(http.StatusOK, "Wellcome to the KUDO CMS - API Gateway!")
}
