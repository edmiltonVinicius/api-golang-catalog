package config

import (
	"errors"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

var validationMessages = map[string]string{
	"required": "campo obrigatório",
	"min":      "tamanho mínimo não atingido",
	"max":      "tamanho máximo excedido",
	"gt":       "deve ser maior que zero",
	"lt":       "deve ser menor que zero",
	"eq":       "deve ser igual a zero",
	"ne":       "deve ser diferente de zero",
	"oneof":    "deve ser um dos valores permitidos",
	"notoneof": "deve ser um dos valores permitidos",
	"len":      "deve ter o tamanho correto",
	"minlen":   "deve ter o tamanho mínimo correto",
	"maxlen":   "deve ter o tamanho máximo correto",
}

func ValidationErrors(err error) []ValidationError {
	var validationErrs validator.ValidationErrors
	if !errors.As(err, &validationErrs) {
		return nil
	}

	response := make([]ValidationError, 0, len(validationErrs))

	for _, e := range validationErrs {
		log.Println("Error validating request:", e)
		response = append(response, ValidationError{
			Field:   strings.ToLower(e.Field()),
			Message: validationMessages[e.Tag()],
		})
	}

	return response
}
