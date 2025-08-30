package myApp

import (
	"fmt"
	"net/http"
	"os"

	"github.com/alamgir-ahosain/e-commerce-project/config"
	"github.com/alamgir-ahosain/e-commerce-project/internal/api/routes"
	"github.com/alamgir-ahosain/e-commerce-project/internal/middleware"
	"github.com/alamgir-ahosain/e-commerce-project/internal/middleware/global"
)

func Serve() {

	//load config
	cnf := config.GetConfig()
	address := ":" + fmt.Sprint(cnf.HttpPort)

	// Middlewaere
	manager := middleware.NewManager()
	global.GlobalMiddleware(manager)

	mux := http.NewServeMux()          //router
	wrappedMux := manager.WrapMux(mux) //warp with Global Middleware
	routes.RegisterRoutes(mux, manager)

	fmt.Println("Server running on port", address)
	// err := http.ListenAndServe(":8080", wrappedMux)
	err := http.ListenAndServe(address, wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}
