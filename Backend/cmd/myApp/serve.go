package myApp

import (
	"fmt"
	"os"

	"github.com/alamgir-ahosain/e-commerce-project/config"
	"github.com/alamgir-ahosain/e-commerce-project/infrastructure/database"
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
	fmt.Println("vesion:", cnf.Version, " ServiceName: ", cnf.ServiceName, " HttpPort: ", cnf.HttpPort)

	dbCon, err := database.NewConnection(cnf.DB)
	fmt.Println("DB connection String: ", dbCon)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//migrate Database
	err = database.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// product
	middlewares := middleware.NewMiddlewares(cnf)
	productRepo := repository.NewProductRepo(dbCon)
	productHandler := product.NewHandler(middlewares, productRepo)

	//user
	userRepo := repository.NewUserRepo(dbCon)
	userHandler := user.NewHandler(cnf, userRepo)

	//Review Section
	reviewHandler := review.NewHandler()

	server := internal.NewServer(cnf, productHandler, userHandler, reviewHandler)
	server.Start()
}
