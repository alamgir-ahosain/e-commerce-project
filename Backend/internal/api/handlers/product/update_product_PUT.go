package product

import (
	"net/http"

	"github.com/alamgir-ahosain/e-commerce-project/internal/services"
	"github.com/alamgir-ahosain/e-commerce-project/internal/util"
)

func (h *Handler) UpdateProductByIdPUT(w http.ResponseWriter, r *http.Request) {
	id, err := services.GetID(w, r)
	if err != nil {
		util.SendError(w, r, http.StatusBadRequest, err)
		return
	}
	services.UpdateProductByIdPutFunc(w, r, id)
}
