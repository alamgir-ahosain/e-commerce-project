package global

import "github.com/alamgir-ahosain/e-commerce-project/internal/middleware"

func GlobalMiddleware(manager *middleware.Manager) {
	manager.Use(middleware.Logger)

}
