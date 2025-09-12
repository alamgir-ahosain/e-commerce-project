package middleware

import "github.com/alamgir-ahosain/e-commerce-project/config"

type Middlewares struct {
	cnf *config.Config
}

func NewMiddlewares(cnf *config.Config) *Middlewares {
	return &Middlewares{cnf: cnf}
}
