package product

import (
	"net/http"

	"github.com/alamgir-ahosain/e-commerce-project/internal/services"
	"github.com/alamgir-ahosain/e-commerce-project/internal/util"
)

// GET->only header
func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	services.HandleCORSFunc(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	productList, err := h.productRepo.GetProductByIdFunc()
	if err != nil {
		// http.Error(w, "Error showed Product", 400)
		util.SendError(w, http.StatusBadRequest, "Error Product")
		return
	}

	services.MakeJSONFormatThreeFunc(w, http.StatusOK, productList)
	// w.WriteHeader(200)
	// encoder := json.NewEncoder(w)
	// encoder.Encode(productList)
}
