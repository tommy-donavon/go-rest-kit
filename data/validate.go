package data

import (
	"regexp"

	"github.com/go-playground/validator"
)

type Validation struct {
	validate *validator.Validate
}

type ValidationOption struct {
	Name      string
	Operation validator.Func
}

func NewValidator(validationOptions ...ValidationOption) *Validation {
	validate := validator.New()
	for _, v := range validationOptions {
		validate.RegisterValidation(v.Name, v.Operation)
	}
	return &Validation{validate}
}

func NewValidatorFunc(regex string) validator.Func {
	return func(fl validator.FieldLevel) bool {
		re := regexp.MustCompile(regex)
		return re.Match([]byte(fl.Field().String()))
	}
}

func (v *Validation) Validate(s interface{}) error {
	return v.validate.Struct(s)
}
