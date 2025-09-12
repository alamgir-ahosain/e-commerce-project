package product

import (
	"net/http"

	"github.com/alamgir-ahosain/e-commerce-project/internal/api/handlers"
	"github.com/alamgir-ahosain/e-commerce-project/internal/middleware"
)

func (h *Handler) RegisterRoutes(r *http.ServeMux, manager *middleware.Manager) {

	r.Handle("GET /alamgir",
		manager.With(
			http.HandlerFunc(handlers.Test),
		),
	)

	//mux.HandleFunc("/products", getProducts)
	//r.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	r.Handle("GET /products",
		manager.With(
			http.HandlerFunc(h.GetProducts),
			middleware.FirstMiddleware,
			middleware.SecondMiddleware,
		),
	)

	r.Handle("POST /products",
		manager.With(
			http.HandlerFunc(h.CreateProducts),
			h.middlewares.AuthenticateJwt,
			middleware.FirstMiddleware,
			middleware.SecondMiddleware,
		),
	)

	//Get product by id
	r.Handle("GET /products/{id}",
		manager.With(
			http.HandlerFunc(h.GetProductById),
			middleware.FirstMiddleware,
			middleware.SecondMiddleware,
		),
	)

	//Delete product by id
	r.Handle("DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(h.DeleteProductById),
			h.middlewares.AuthenticateJwt,
			middleware.FirstMiddleware,
			middleware.SecondMiddleware,
		),
	)

	//Update product by id
	r.Handle("PUT /products/{id}",
		manager.With(
			http.HandlerFunc(h.UpdateProductById),
			h.middlewares.AuthenticateJwt,
			middleware.FirstMiddleware,
			middleware.SecondMiddleware,
		),
	)
	//Update product by id
	r.Handle("PATCH /products/{id}",
		manager.With(
			http.HandlerFunc(h.UpdateProductByIdPUT),
			h.middlewares.AuthenticateJwt,
			middleware.FirstMiddleware,
			middleware.SecondMiddleware,
		),
	)

}
