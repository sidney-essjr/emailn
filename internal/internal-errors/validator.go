package internalerrors

import (
	"github.com/go-playground/validator/v10"
)

func ValidateStruct[T any](obj T) map[string]string {
	validate := validator.New()
	err := validate.Struct(obj)

	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errors := make(map[string]string)
			for _, fieldError := range validationErrors {
				errors[fieldError.StructField()] = fieldError.Error()
			}
			return errors
		}
		return map[string]string{"error": err.Error()}
	}

	return nil
}
