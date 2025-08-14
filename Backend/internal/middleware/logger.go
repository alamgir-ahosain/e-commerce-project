package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// log.Println("Middleware : Print First")
		next.ServeHTTP(w, r)
		// log.Println("Middleware : Print Last")
		log.Println(r.Method, r.URL.Path, time.Since(start))

	}) 
}
