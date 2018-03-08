package app

import (
	"net/http"
)

// MiddlewareHandler ..
type MiddlewareHandler struct{}

// Middlewares is factory method for MiddlewareHandler
func (app *Kernel) Middlewares() *MiddlewareHandler {
	return &MiddlewareHandler{}
}

// ACustomMiddleware is an example
func (m *MiddlewareHandler) ACustomMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(w, r)
}

// AdminMiddleware for the admin user
func (m *MiddlewareHandler) AdminMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(w, r)
}
