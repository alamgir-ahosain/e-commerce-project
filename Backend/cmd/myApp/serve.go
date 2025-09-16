package myApp

import (
	"fmt"

	"github.com/alamgir-ahosain/e-commerce-project/config"
	"github.com/alamgir-ahosain/e-commerce-project/internal"
	"github.com/alamgir-ahosain/e-commerce-project/internal/api/handlers/product"
	"github.com/alamgir-ahosain/e-commerce-project/internal/api/handlers/review"
	"github.com/alamgir-ahosain/e-commerce-project/internal/api/handlers/user"
	"github.com/alamgir-ahosain/e-commerce-project/internal/middleware"
	"github.com/alamgir-ahosain/e-commerce-project/internal/repository"
)

func Serve() {

	//load config
	cnf := config.GetConfig()
	fmt.Println(cnf.Version)
	fmt.Println(cnf.ServiceName)
	fmt.Println(cnf.HttpPort)

	// product
	middlewares := middleware.NewMiddlewares(cnf)
	productRepo := repository.NewProductRepo()
	productHandler := product.NewHandler(middlewares, productRepo)

	//user
	userRepo := repository.NewUserRepo()
	userHandler := user.NewHandler(cnf, userRepo)

	//Review Section
	reviewHandler := review.NewHandler()

	server := internal.NewServer(cnf, productHandler, userHandler, reviewHandler)
	server.Start()
}
