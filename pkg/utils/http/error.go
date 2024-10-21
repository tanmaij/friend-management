package http

import (
	"encoding/json"
	"net/http"
)

// Error represents custom error
type Error struct {
	Status  int    `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// WriteToHttpResponseWriter writes a custom error into http.ResponseWriter
func (e Error) WriteToHttpResponseWriter(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Status)

	decodedMessage, err := json.Marshal(e)
	if err != nil {
		return err
	}

	w.Write(decodedMessage)
	return nil
}
