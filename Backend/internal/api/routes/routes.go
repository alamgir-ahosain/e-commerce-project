package routes

import (
	"net/http"

	"github.com/alamgir-ahosain/e-commerce-project/internal/api/handlers"
	"github.com/alamgir-ahosain/e-commerce-project/internal/middleware"
)

func RegisterRoutes(r *http.ServeMux, manager *middleware.Manager) {

	r.Handle("GET /alamgir",
		manager.With(
			http.HandlerFunc(handlers.Test),
		),
	)

	//mux.HandleFunc("/products", getProducts)
	//r.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	r.Handle("GET /products",
		manager.With(
			http.HandlerFunc(handlers.GetProducts),
			middleware.FirstMiddleware,
			middleware.SecondMiddleware,
		),
	)

	r.Handle("POST /products",
		manager.With(
			http.HandlerFunc(handlers.CreateProducts),
			middleware.FirstMiddleware,
			middleware.SecondMiddleware,
		),
	)

	//Get product by id
	r.Handle("GET /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.GetProductById),
			middleware.FirstMiddleware,
			middleware.SecondMiddleware,
		),
	)
	//Delete product by id
	r.Handle("DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.DeleteProductById),
			middleware.FirstMiddleware,
			middleware.SecondMiddleware,
		),
	)

}
