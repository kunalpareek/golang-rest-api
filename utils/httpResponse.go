package utils

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON returns an http response with custom payload
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	//fmt.Println("Going to send respones")
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// RespondWithError returns an http response with error message
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJSON(w, code, map[string]string{"error": msg})
}
