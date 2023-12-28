package validators

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

func LoadValidators() error {
	Validate = validator.New()

	err := Validate.RegisterValidation("name", NameValidator)
	if err != nil {
		return err
	}

	err = Validate.RegisterValidation("password", PasswordValidator)
	if err != nil {
		return err
	}

	return nil
}
