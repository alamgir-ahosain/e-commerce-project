package myApp

import (
	"fmt"
	"net/http"

	"github.com/alamgir-ahosain/e-commerce-project/internal/api/handlers"
	"github.com/alamgir-ahosain/e-commerce-project/internal/api/routes"
)

func Serve() {
	mux := http.NewServeMux() //router
	routes.RegisterRoutes(mux)

	globalRouter := handlers.GlobalRouter(mux) //handle CORS and OPTIONS method
	fmt.Println("Server running on port :8080")
	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
