package review

import (
	"net/http"

	"github.com/alamgir-ahosain/e-commerce-project/internal/middleware"
)

func (h *Handler) RegisterRoutes(r *http.ServeMux, manager *middleware.Manager) {
	// Create User
	r.Handle("GET /review",
		manager.With(
			http.HandlerFunc(h.GetReview),
			middleware.FirstMiddleware,
			middleware.SecondMiddleware,
		),
	)

}
