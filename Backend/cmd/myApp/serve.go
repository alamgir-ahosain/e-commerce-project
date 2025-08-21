package myApp

import (
	"fmt"
	"net/http"

	"github.com/alamgir-ahosain/e-commerce-project/internal/api/routes"
	"github.com/alamgir-ahosain/e-commerce-project/internal/middleware"
	"github.com/alamgir-ahosain/e-commerce-project/internal/middleware/global"
)

func Serve() {
	mux := http.NewServeMux() //router

	// Middlewaere
	manager := middleware.NewManager()
	global.GlobalMiddleware(manager)

	routes.RegisterRoutes(mux, manager)
	globalRouter := middleware.CORSwithPreflight(mux) //handle CORS and OPTIONS method
	fmt.Println("Server running on port :8080")
	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
