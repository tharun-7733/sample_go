package shared

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func InternalErrorHandler(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	response := ErrorResponse{
		Error: "Internal server error",
		Code:  http.StatusInternalServerError,
	}

	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(response)
}

func RequestErrorHandler(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	response := ErrorResponse{
		Error: err.Error(),
		Code:  http.StatusUnauthorized,
	}

	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(response)
}