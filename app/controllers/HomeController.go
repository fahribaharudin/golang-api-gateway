package controllers

import (
	"net/http"
)

// HomeController handler to some endpoint
type HomeController struct {
}

// LandingEndpointHandler to GET:/
func (controller *HomeController) LandingEndpointHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Kudo CMS - API Gateway!"))
}
