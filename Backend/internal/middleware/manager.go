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

// Receiver function,wrapper local middleware(First,second)
func (mn *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	n := next
	for _, middleware := range middlewares {
		n = middleware(n)
		//n =FirstMiddleware(GetProducts)  -> n= SecondMiddleware(FirstMiddleware(GetProducts))
	}
	return n
}

// wrapper Global Middleware with Mux
func (mn *Manager) WrapMux(next http.Handler) http.Handler {
	n := next // [preflight,cors,logger]   n=logger(cors(preflight(mux)))
	for _, middleware := range mn.globalMiddlewares {
		n = middleware(n)
	}
	return n
}

// retrun new Middleware object
func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}
