package httpapi

import "github.com/go-playground/validator/v10"

func ValidationErrors(err error) map[string]string {
	errors := make(map[string]string)

	for _, e := range err.(validator.ValidationErrors) {
		errors[e.Field()] = e.Tag()
	}

	return errors
}
