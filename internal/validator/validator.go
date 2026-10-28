// Package validator provides a validator v10
package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Validator struct
type Validator struct {
	validate *validator.Validate
}

// New creates a new Validator
func New() *Validator {
	return &Validator{
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}

// ValidationError struct
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Validate validates a struct and returns a slice of ValidationError
func (v *Validator) Validate(s interface{}) []ValidationError {
	err := v.validate.Struct(s)
	if err != nil {
		return nil
	}

	var errors []ValidationError

	validationErrors, ok := err.(validator.ValidationErrors)

	if !ok {
		return nil
	}

	for _, e := range validationErrors {
		errors = append(errors, ValidationError{
			Field:   strings.ToLower(e.Field()),
			Message: translateTag(e),
		})
	}
	return errors

}

func translateTag(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "este campo es obligatorio"
	case "email":
		return "este campo debe ser un correo electrónico válido"
	case "min":
		return fmt.Sprintf("este campo debe tener como mínimo %s caracteres", e.Param())
	case "max":
		return fmt.Sprintf("este campo debe tener como máximo %s caracteres", e.Param())
	default:
		return "valor inválido"
	}
}
