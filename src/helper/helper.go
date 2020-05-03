package helper

import (
	"project/src/helper/response"
	"project/src/helper/validator"
	"project/src/helper/viper"
)

// Helper ...
type Helper struct {
	Response  response.Method
	Config    viper.Method
	Validator validator.Method
}

// NewHelper ...
func NewHelper() Helper {
	return Helper{
		Response:  response.NewResponse(),
		Config:    viper.NewViper(),
		Validator: validator.NewValidator(),
	}
}
