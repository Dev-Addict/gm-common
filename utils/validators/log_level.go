package validators

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func LegLevelValidator(fl validator.FieldLevel) bool {
	if fl.Field().String() == "" {
		return true
	}

	r := regexp.MustCompile(`^(debug|info|warn|error|fatal|panic)$`)

	return r.MatchString(fl.Field().String())
}
