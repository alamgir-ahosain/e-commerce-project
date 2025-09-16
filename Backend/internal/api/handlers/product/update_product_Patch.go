package product

import (
	"encoding/json"
	"net/http"

	"github.com/alamgir-ahosain/e-commerce-project/internal/services"
	"github.com/alamgir-ahosain/e-commerce-project/internal/util"
)

func (h *Handler) UpdateProductByIdPatch(w http.ResponseWriter, r *http.Request) {
	id, err := services.GetID(w, r)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, err)
		return
	}

	// services.UpdateProductByIdPatchFunc(w, r, id)
	var patchData map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&patchData)
	if err != nil {
		http.Error(w, "Give the valid JSON format", http.StatusBadRequest)
		return
	}
	updatedProduct, err := h.productRepo.UpdateProductByIdPatchFunc(patchData, id)
	if err != nil {
		util.SendError(w, http.StatusNotFound, err)
		return
	}

	services.MakeJSONFormatThreeFunc(w, http.StatusOK, updatedProduct)

}
