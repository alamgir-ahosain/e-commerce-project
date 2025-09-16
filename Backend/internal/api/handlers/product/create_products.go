package product

import (
	"encoding/json"
	"net/http"

	"github.com/alamgir-ahosain/e-commerce-project/internal/models"
	"github.com/alamgir-ahosain/e-commerce-project/internal/services"
	"github.com/alamgir-ahosain/e-commerce-project/internal/util"
)

// POST->header and body
func (h *Handler) CreateProducts(w http.ResponseWriter, r *http.Request) {

	services.HandleCORSFunc(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	/*
		1. take body information (description,price,title,imageUrl) from r.Body()
		2.create an instance using Product struct with the body information
		3.append the instance into productList
	*/

	var newProduct models.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Give  the valid json format", http.StatusBadRequest)
		return
	}

	//
	createdProduct, err := h.productRepo.CreateProductFunc(newProduct)
	if err != nil {
		// http.Error(w, "Error Creating Product", 400)
		util.SendError(w, http.StatusBadRequest, "Error Creating Product")
		return
	}
	services.MakeJSONFormatThreeFunc(w, http.StatusCreated, createdProduct)

}
