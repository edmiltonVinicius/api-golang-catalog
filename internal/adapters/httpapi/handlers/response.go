package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/edmiltonVinicius/go-api-catalog/internal/config"
)

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func writeError(w http.ResponseWriter, status int, err error) {
	writeJSON(w, status, map[string]string{
		"error": err.Error(),
	})
}

func writeValidationErrors(w http.ResponseWriter, status int, errors []config.ValidationError) {
	writeJSON(w, status, map[string][]config.ValidationError{
		"errors": errors,
	})
}
