package middleware

import (
	"net/http"
)

type Middleware func(next http.Handler) http.Handler
type Manager struct {
	globalMiddlewares []Middleware //list of middleware
}

// builder patter: add middleware
func (mn *Manager) Use(middleware ...Middleware) {
	mn.globalMiddlewares = append(mn.globalMiddlewares, middleware...)

}

// Receiver function
func (mn *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	n := next

	//call middleware
	for _, middleware := range middlewares {
		n = middleware(n)
	}

	//call globalMiddleware
	for _, globalMiddleware := range mn.globalMiddlewares {
		n = globalMiddleware((n))
	}

	return n
}

// retrun new Middleware object
func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}
