package product

import (
	"encoding/json"
	"net/http"

	"github.com/alamgir-ahosain/e-commerce-project/internal/models"
	"github.com/alamgir-ahosain/e-commerce-project/internal/services"
	"github.com/alamgir-ahosain/e-commerce-project/internal/util"
)

func (h *Handler) UpdateProductById(w http.ResponseWriter, r *http.Request) {

	id, err := services.GetID(w, r)
	if err != nil {
		return
	}

	var newProduct models.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Give  the valid json format", http.StatusBadRequest)
		return
	}

	product, err := h.productRepo.UpdateProductByIdFunc(newProduct, id)

	if err != nil {
		util.SendError(w, http.StatusBadRequest, err)
		return
	}

	services.MakeJSONFormatThreeFunc(w, http.StatusOK, product)

}
