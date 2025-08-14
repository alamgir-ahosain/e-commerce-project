package middleware

import (
	"log"
	"net/http"
)

func FirstMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		log.Println("First Middleware ")
		next.ServeHTTP(w,r)
	})
}