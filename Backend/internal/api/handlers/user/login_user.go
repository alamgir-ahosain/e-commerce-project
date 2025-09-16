package user

import (
	"encoding/json"
	"net/http"

	"github.com/alamgir-ahosain/e-commerce-project/internal/models"
	"github.com/alamgir-ahosain/e-commerce-project/internal/services"
	"github.com/alamgir-ahosain/e-commerce-project/internal/util"
)

// POST->header and body
func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {

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

	var reqLogin models.RequestLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)
	if err != nil {
		http.Error(w, "Give  the valid json format", http.StatusBadRequest)
		return

	}

	user, err := h.userRepo.LoginUserFunc(reqLogin)
	if err != nil {
		util.SendError(w, http.StatusNotFound, err)
		return
	}

	accessToken, err := util.CreateJWT(h.config.JwtSecretKey, util.Payload{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		IsShopOwner: user.IsShopOwner,
	})
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, err)
		return
	}
	// services.MakeJSONFormatThreeFunc(w, http.StatusOK, pr)
	services.MakeJSONFormatThreeFunc(w, http.StatusOK, accessToken)

}
