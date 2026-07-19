package validators

import "github.com/go-playground/validator/v10"

var validate validator.Validate

func NewValidator() validator.Validate{
	return *validator.New(
		validator.WithRequiredStructEnabled(),
	)
}

func init() {
	validate = NewValidator()
}