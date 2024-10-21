package http

import (
	"encoding/json"
	"net/http"
)

// WriteString writes the response with string body to the given http.Writer with success status
func WriteString(v string, httpStatus int, w http.ResponseWriter) {
	w.WriteHeader(httpStatus)
	w.Write([]byte(v))
}

// WriteJsonData writes the response with Json body to the given http.Writer with success status
func WriteJsonData(v any, httpStatus int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	decodedData, _ := json.Marshal(v)
	w.Write(decodedData)
}
