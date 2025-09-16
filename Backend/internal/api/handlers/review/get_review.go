package review

import (
	"net/http"

	"github.com/alamgir-ahosain/e-commerce-project/internal/services"
)

// POST->header and body
func (h *Handler) GetReview(w http.ResponseWriter, r *http.Request) {

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

	// services.CreateProductFunc(w, r)
	//services.CreateUserFunc(w, r)

}
