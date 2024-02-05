package initializers

import "github.com/go-playground/validator/v10"

var VD *validator.Validate

func Validator() {
	VD = validator.New()
}
