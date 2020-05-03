package validator

import "gopkg.in/go-playground/validator.v9"

type validatorHelper struct {
	validate *validator.Validate
}

// Method ...
type Method interface {
	Validate(i interface{}) error
	Validator() *validatorHelper
}

// NewValidator ...
func NewValidator() Method {
	return &validatorHelper{
		validate: validator.New(),
	}
}

func (t *validatorHelper) Validate(i interface{}) error {
	return t.validate.Struct(i)
}

func (t *validatorHelper) Validator() *validatorHelper {
	return t
}
