package util

import (
	"encoding/json"
	"net/http"
)

func SendError(w http.ResponseWriter, r *http.Request, statusCode int, msg interface{}) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(msg)

}
