package user

import (
	"encoding/json"
	"net/http"

	"github.com/alamgir-ahosain/e-commerce-project/internal/models"
	"github.com/alamgir-ahosain/e-commerce-project/internal/services"
	"github.com/alamgir-ahosain/e-commerce-project/internal/util"
)

// POST->header and body
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	services.HandleCORSFunc(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
	/*
		1. take body information (description,price,title,imageUrl) from r.Body()
		2.create an instance using Product struct with the body information
		3.append the instance into productList
	*/

	var newUser models.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Give the valid json format")
		return
	}

	user, err := h.userRepo.CreateUserFunc(newUser)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Error Creating User")
		return
	}
	services.MakeJSONFormatThreeFunc(w, http.StatusCreated, user)

}
