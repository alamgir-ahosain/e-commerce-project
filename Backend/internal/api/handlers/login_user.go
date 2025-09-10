package handlers

import (
	"net/http"

	"github.com/alamgir-ahosain/e-commerce-project/config"
	"github.com/alamgir-ahosain/e-commerce-project/internal/services"
	"github.com/alamgir-ahosain/e-commerce-project/internal/util"
)

// POST->header and body
func LoginUser(w http.ResponseWriter, r *http.Request) {

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

	user, err := services.LoginUserFunc(w, r)
	if err != nil {
		util.SendError(w, r, http.StatusNotFound, err)
		return
	}

	cnf := config.GetConfig()
	accessToken, err := util.CreateJWT(cnf.JwtSecretKey, util.Payload{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		IsShopOwner: user.IsShopOwner,
	})
	if err != nil {
		util.SendError(w, r, http.StatusInternalServerError, err)
		return
	}
	// services.MakeJSONFormatThreeFunc(w, http.StatusOK, pr)
	services.MakeJSONFormatThreeFunc(w, http.StatusOK, accessToken)

}
