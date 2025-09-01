package handlers

import (
	"net/http"

	"github.com/alamgir-ahosain/e-commerce-project/internal/services"
	"github.com/alamgir-ahosain/e-commerce-project/internal/util"
)

func DeleteProductById(w http.ResponseWriter, r *http.Request) {
	id, err := services.GetID(w, r)
	if err != nil {
		util.SendError(w, r, http.StatusNotFound, err)
		return
	}
	_, err = services.DeleteProductByIdFunc(id)
	if err != nil {
		util.SendError(w, r, http.StatusBadRequest, err)
		return
	}
}
