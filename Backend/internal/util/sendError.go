package util

import (
	"encoding/json"
	"net/http"
)

// func SendError(w http.ResponseWriter, statusCode int, msg interface{}) {
// 	w.WriteHeader(statusCode)
// 	encoder := json.NewEncoder(w)
// 	encoder.Encode(msg)

// }

func SendError(w http.ResponseWriter, statusCode int, msg interface{}) {
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	var errMsg string

	// Determine the error message based on the type
	switch v := msg.(type) {
	case error:
		errMsg = v.Error()
	case string:
		errMsg = v
	default:
		errMsg = "unknown error"
	}

	json.NewEncoder(w).Encode(map[string]string{
		"error": errMsg,
	})
}
