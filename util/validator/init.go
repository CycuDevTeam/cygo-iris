package validator

import (
	"github.com/go-playground/validator/v10"
)

func NewValidator() *validator.Validate {
	v := validator.New()
	InitRegexValidator(v)
	return v
}

// Add costum validate rule here.
func InitRegexValidator(v *validator.Validate) {
	_ = v.RegisterValidation("is_username", isUsername)
}
