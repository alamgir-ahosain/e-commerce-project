package routes

import (
	"net/http"
	"github.com/alamgir-ahosain/e-commerce-project/internal/api/handlers"
)

func RegisterRoutes(r *http.ServeMux) {
	//mux.HandleFunc("/products", getProducts)  
	r.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	r.Handle("POST /create-products", http.HandlerFunc(handlers.CreateProduct))

}
