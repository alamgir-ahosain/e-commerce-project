package handlers

import (
	"net/http"

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

	pr, err := services.LoginUserFunc(w, r)
	if err != nil {
		util.SendError(w, r, http.StatusNotFound, err)
		return
	}
	services.MakeJSONFormatThreeFunc(w, http.StatusOK, pr)

}
