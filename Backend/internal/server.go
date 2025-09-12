package internal

import (
	"fmt"
	"net/http"
	"os"

	"github.com/alamgir-ahosain/e-commerce-project/config"
	"github.com/alamgir-ahosain/e-commerce-project/internal/api/handlers/product"
	"github.com/alamgir-ahosain/e-commerce-project/internal/api/handlers/review"
	"github.com/alamgir-ahosain/e-commerce-project/internal/api/handlers/user"
	"github.com/alamgir-ahosain/e-commerce-project/internal/middleware"
	"github.com/alamgir-ahosain/e-commerce-project/internal/middleware/global"
)

type Server struct {
	cnf            *config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
	reviewHandler  *review.Handler
}

func NewServer(cnf *config.Config, productHandler *product.Handler, userHandler *user.Handler, reviewHandler *review.Handler) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
		reviewHandler:  reviewHandler,
	}
}

func (server *Server) Start() {
	address := ":" + fmt.Sprint(server.cnf.HttpPort)

	// Middlewaere
	manager := middleware.NewManager()
	global.GlobalMiddleware(manager)

	mux := http.NewServeMux()          //router
	wrappedMux := manager.WrapMux(mux) //warp with Global Middleware

	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)
	server.reviewHandler.RegisterRoutes(mux, manager)

	fmt.Println("Server running on port", address)
	// err := http.ListenAndServe(":8080", wrappedMux)
	err := http.ListenAndServe(address, wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}
