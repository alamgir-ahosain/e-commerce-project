package product

import (
	"net/http"

	"github.com/alamgir-ahosain/e-commerce-project/internal/services"
)

// GET->only header
func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	services.HandleCORSFunc(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(200)
		return
	}

	services.MakeJSONFormatFunc(w, 200)
}
