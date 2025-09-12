package user

import (
	"net/http"

	"github.com/alamgir-ahosain/e-commerce-project/internal/middleware"
)

func (h *Handler) RegisterRoutes(r *http.ServeMux, manager *middleware.Manager) {
	// Create User
	r.Handle("POST /users",
		manager.With(
			http.HandlerFunc(h.CreateUser),
			middleware.FirstMiddleware,
			middleware.SecondMiddleware,
		),
	)

	//Login User
	r.Handle("POST /users/login",
		manager.With(
			http.HandlerFunc(h.LoginUser),
			middleware.FirstMiddleware,
			middleware.SecondMiddleware,
		),
	)
}
