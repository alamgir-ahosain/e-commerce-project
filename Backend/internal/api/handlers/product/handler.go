package product

import "github.com/alamgir-ahosain/e-commerce-project/internal/middleware"

type Handler struct {
	middlewares *middleware.Middlewares
}

func NewHandler(middlewares *middleware.Middlewares) *Handler {
	return &Handler{middlewares: middlewares}
}
