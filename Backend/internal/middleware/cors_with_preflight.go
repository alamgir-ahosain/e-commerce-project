package middleware

import "net/http"

// for all request: First handle CORS and then Preflight request(OPTIONS method)
func CORSwithPreflight(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                                  //set header: (*) anyone can access the responce
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Alamgir-CustomHeader") //allow content type
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS") //allow content type
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}
		next.ServeHTTP(w, r)
	})
}
