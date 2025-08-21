package middleware

import (
	"log"
	"net/http"
)

// for all request: First handle CORS and then Preflight request(OPTIONS method)
func Cors(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("CORS middleware")
		
		w.Header().Set("Access-Control-Allow-Origin", "*")                                  //set header: (*) anyone can access the responce
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Alamgir-CustomHeader") //allow content type
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS") //allow content type
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
