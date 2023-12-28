package validators

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func NameValidator(fl validator.FieldLevel) bool {
	if fl.Field().String() == "" {
		return true
	}

	r := regexp.MustCompile(`^([A-Z][a-z]*\s?)+$`)

	return r.MatchString(fl.Field().String())
}
