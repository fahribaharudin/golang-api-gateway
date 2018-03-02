package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIGatewayController is the http proxy handler
type APIGatewayController struct {
}

// GetHandler handle GET
func (controller *APIGatewayController) GetHandler(c *gin.Context) {
	c.String(http.StatusOK, "Well this is GET ?")
}

// PostHandler handle POST
func (controller *APIGatewayController) PostHandler(c *gin.Context) {
	c.String(http.StatusOK, "Well this is POST ?")
}

// PutHandler handle PUT
func (controller *APIGatewayController) PutHandler(c *gin.Context) {
	c.String(http.StatusOK, "Well this is PUT ?")
}

// PatchHandler handle PATCH
func (controller *APIGatewayController) PatchHandler(c *gin.Context) {
	c.String(http.StatusOK, "Well this is PATCH ?")
}

// DeleteHandler handle DELETE
func (controller *APIGatewayController) DeleteHandler(c *gin.Context) {
	c.String(http.StatusOK, "Well this is DELETE ?")
}
