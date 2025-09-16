package product

import (
	"github.com/alamgir-ahosain/e-commerce-project/internal/middleware"
	"github.com/alamgir-ahosain/e-commerce-project/internal/repository"
)

type Handler struct {
	middlewares *middleware.Middlewares
	productRepo repository.ProductRepo
}

func NewHandler(middlewares *middleware.Middlewares, productRepo repository.ProductRepo) *Handler {
	return &Handler{
		middlewares: middlewares,
		productRepo: productRepo,
	}
}
