package validation

import (
	"github.com/go-playground/validator/v10"
)

// Validator the validator service
type Validator struct {
	validator *validator.Validate
}

// NewValidator returns a new Validator
func NewValidator() *Validator {
	return &Validator{validator: validator.New()}
}

// Validate validates the given interface
func (v *Validator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		vErrFields := []*ErrorField{}

		for _, vErr := range err.(validator.ValidationErrors) {
			e := NewErrorField(vErr)
			vErrFields = append(vErrFields, e)
		}

		return &ErrorValidation{
			ErrorFields: vErrFields,
		}
	}

	return nil
}
