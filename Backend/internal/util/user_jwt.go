package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

// JWT Header
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

// JWT Playload
type Payload struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner bool
}

/*
			steps for JWT creation
	1. Header -> Base64 (b64Header)
	2. Payload -> Base64 (b64Payload)
	3. Combine: headerPayload := b64Header + "." + b64Payload
	4. Signature: HMAC-SHA256(secretKey, headerPayload)
	5. Encode signature -> Base64 (hash)
	6. JWT:  jwt := b64Header + "." + b64Payload + "." + hash



*/

// Create JWT
func CreateJWT(secretKey string, data Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	// Secret Key : Encode to byte
	byteArrSecret := []byte(secretKey)

	// Header : Encode to byte
	byteArrHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	// Header: Encode to base64
	headerBase64 := base64String(byteArrHeader)

	// Payload : convert to byte
	byteArrPayload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// Payload: Encode payload to base64
	payloadBase64 := base64String(byteArrPayload)

	// Singnature: Create HMAC-SHA256 signature
	message := headerBase64 + "." + payloadBase64
	h := hmac.New(sha256.New, byteArrSecret)
	h.Write([]byte(message))
	signature := h.Sum(nil)

	// Encode signature to base64
	signatureBase64 := base64String(signature)

	// Combine header, payload, and signature to form JWT
	jwt := headerBase64 + "." + payloadBase64 + "." + signatureBase64
	return jwt, nil
}

func base64String(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
