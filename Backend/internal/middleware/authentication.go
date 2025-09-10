package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/alamgir-ahosain/e-commerce-project/config"
)

func AuthenticateJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log.Println("Authintication Middleware ")
		//JWT
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		//spilt into Bearer and secret
		headerArr := strings.Split(header, " ")
		accessToken := headerArr[1]
		if len(headerArr) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(accessToken, ".")
		if len(tokenParts) != 3 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		jwtHeader := tokenParts[0]
		playload := tokenParts[1]
		signature := tokenParts[2]

		// Singnature: Create HMAC-SHA256 signature
		message := jwtHeader + "." + playload

		byteArrSecret := []byte(config.GetConfig().JwtSecretKey)
		byteArrMsg := []byte(message)

		h := hmac.New(sha256.New, []byte(byteArrSecret))
		h.Write([]byte(byteArrMsg))
		hash := h.Sum(nil)

		// Encode signature to base64
		newSignature := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(hash)
		if signature != newSignature {
			http.Error(w, "Unauthorized,hacker found", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
