package myApp

import (
	"fmt"
	"net/http"

	"github.com/alamgir-ahosain/e-commerce-project/internal/api/routes"
	"github.com/alamgir-ahosain/e-commerce-project/internal/middleware"
	"github.com/alamgir-ahosain/e-commerce-project/internal/middleware/global"
)

func Serve() {

	// Middlewaere
	manager := middleware.NewManager()
	global.GlobalMiddleware(manager)


	mux := http.NewServeMux() //router
	wrappedMux := manager.WrapMux(mux) //warp with Global Middleware
	routes.RegisterRoutes(mux, manager)

	fmt.Println("Server running on port :8080")
	err := http.ListenAndServe(":8080", wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
