package routes

import (
	"net/http"

	"github.com/alamgir-ahosain/e-commerce-project/internal/api/handlers"
	"github.com/alamgir-ahosain/e-commerce-project/internal/middleware"
)

func RegisterRoutes(r *http.ServeMux) {
	
	
	r.Handle("GET /alamgir", middleware.Logger( http.HandlerFunc(handlers.Test)))

	//mux.HandleFunc("/products", getProducts)
	r.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	r.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))

}
