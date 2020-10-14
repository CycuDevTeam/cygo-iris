package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

const (
	usernameString = "^[a-zA-Z0-9_-]+$"
)

var (
	usernameRegex = regexp.MustCompile(usernameString)
)

func isUsername(fl validator.FieldLevel) bool {
	return usernameRegex.MatchString(fl.Field().String())
}
