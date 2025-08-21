package middleware

import (
	"log"
	"net/http"
)

// for all request: First handle CORS and then Preflight request(OPTIONS method)
func Preflight(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Preflight")

		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}
		next.ServeHTTP(w, r)
	})
}
