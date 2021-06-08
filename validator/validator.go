package validator

import "github.com/go-playground/validator"

var validate *validator.Validate

func Validation(model interface{}) []validator.FieldError {
	var errors []validator.FieldError
	validate = validator.New()
	err := validate.Struct(model)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err)
		}
	}
	return errors
}
