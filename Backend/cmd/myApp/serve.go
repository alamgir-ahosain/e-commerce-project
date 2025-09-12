package myApp

import (
	"fmt"

	"github.com/alamgir-ahosain/e-commerce-project/config"
	"github.com/alamgir-ahosain/e-commerce-project/internal"
	"github.com/alamgir-ahosain/e-commerce-project/internal/api/handlers/product"
	"github.com/alamgir-ahosain/e-commerce-project/internal/api/handlers/review"
	"github.com/alamgir-ahosain/e-commerce-project/internal/api/handlers/user"
	"github.com/alamgir-ahosain/e-commerce-project/internal/middleware"
)

func Serve() {

	//load config
	cnf := config.GetConfig()
	fmt.Println(cnf.Version)
	fmt.Println(cnf.ServiceName)
	fmt.Println(cnf.HttpPort)

	mid := middleware.NewMiddlewares(cnf)
	productHandler := product.NewHandler(mid)
	userHandler := user.NewHandler()
	reviewHandler := review.NewHandler()

	server := internal.NewServer(cnf, productHandler, userHandler, reviewHandler)
	server.Start()
}
