package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Error represents custom error
type Error struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Error implements Error method
func (e Error) Error() string {
	return fmt.Sprintf("Status: %d, Code: %s, Desc: %s", e.Status, e.Code, e.Message)
}

// WriteErrorToHttpResponseWriter writes a custom error into http.ResponseWriter
func WriteErrorToHttpResponseWriter(w http.ResponseWriter, e Error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Status)

	encodedData, err := json.Marshal(e)
	if err != nil {
		return err
	}

	w.Write(encodedData)
	return nil
}
