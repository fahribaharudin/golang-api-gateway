package controllers

import (
	"net/http"
)

// APIGatewayController is the http proxy handler
type APIGatewayController struct {
}

// UniversalHandler to handle all request for the API Gateway
func (c *APIGatewayController) UniversalHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Its from universal handler dude, you want to access: " + r.RequestURI + " with: " + r.Method + " right?"))
}
