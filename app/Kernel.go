package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fahribaharudin/api_gateway/app/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// Kernel is the app skeleton
type Kernel struct {
	Router      *mux.Router
	Middleware  *negroni.Negroni
	Controllers Controllers
	HTTPHandler http.Handler
}

// Controllers is a wrapper of all the controllers
type Controllers struct {
	APIGateway *controllers.APIGateway
}

// Bootstrap the app Kernel
func (app *Kernel) Bootstrap() {
	app.Router = mux.NewRouter()
	app.Middleware = negroni.Classic()
	app.Controllers = Controllers{
		APIGateway: &controllers.APIGateway{},
	}
}

// Run the application
func (app *Kernel) Run() {
	if (app.Controllers == Controllers{}) {
		fmt.Println("No controllers defined in the app kernel, You should call the Bootstrap() method first before running the app")
		os.Exit(0)
	}

	// registering all the endpoint handler (routes) to the router
	app.RegisterRoutes()

	// registering the global middleware
	app.Middleware.Use(negroni.HandlerFunc(app.Middlewares().ACustomMiddleware))
	app.Middleware.UseHandler(app.Router)
	app.HTTPHandler = app.Middleware

	// make the server
	server := &http.Server{
		Addr:         ":8000",
		Handler:      app.HTTPHandler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	log.Println("Server listening on port" + server.Addr)
	server.ListenAndServe()
}
