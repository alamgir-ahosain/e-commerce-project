package handlers

import (
	"net/http"

	"github.com/alamgir-ahosain/e-commerce-project/internal/services"
)

func UpdateProductById(w http.ResponseWriter, r *http.Request) {

	id, err := services.GetID(w, r)
	if err != nil {
		return
	}
	services.UpdateProductByIdFunc(w, r, id)

}
