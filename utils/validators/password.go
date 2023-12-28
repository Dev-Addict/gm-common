package validators

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func PasswordValidator(fl validator.FieldLevel) bool {
	if fl.Field().String() == "" {
		return true
	}

	regexps := []string{".{8,}", "[a-z]", "[A-Z]", "[0-9]", "[^\\d\\w]"}

	for _, r := range regexps {
		if !regexp.MustCompile(r).MatchString(fl.Field().String()) {
			return false
		}
	}

	return true
}
