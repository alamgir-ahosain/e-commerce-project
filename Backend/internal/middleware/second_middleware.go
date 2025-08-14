package middleware

import (
	"log"
	"net/http"
)

func SecondMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		log.Println("Second Middleware ")
		next.ServeHTTP(w,r)
	})
}