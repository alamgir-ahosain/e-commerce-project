package product

import (
	"net/http"

	"github.com/alamgir-ahosain/e-commerce-project/internal/services"
	"github.com/alamgir-ahosain/e-commerce-project/internal/util"
)

func (h *Handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	id, err := services.GetID(w, r)
	if err != nil {
		return
	}
	// pr, err := services.GetProductByIdFunc(id)
	pr, err := h.productRepo.GetByIdFunc(id)
	if err != nil {
		util.SendError(w, http.StatusNotFound, err)
		return
	}
	services.MakeJSONFormatThreeFunc(w, http.StatusOK, pr)

}
