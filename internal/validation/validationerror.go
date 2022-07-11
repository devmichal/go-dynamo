package validation

import "fmt"

type ErrorField struct {
	FieldName     string
	FieldValue    interface{}
	ValidationTag string
}

// ErrorValidation the validation error for all the fields
type ErrorValidation struct {
	ErrorFields []*ErrorField
}

// FieldErrorInterface the interface for the filed error
type FieldErrorInterface interface {
	Namespace() string
	Value() interface{}
	Tag() string
}

// NewErrorField creates a new ErrorField
func NewErrorField(ve FieldErrorInterface) *ErrorField {
	return &ErrorField{
		FieldName:     ve.Namespace(),
		FieldValue:    ve.Value(),
		ValidationTag: ve.Tag(),
	}
}

// Error the Error method implementation of the error interface
func (e *ErrorValidation) Error() string {
	errStr := "Validation error:"

	for _, err := range e.ErrorFields {
		fErrStr := fmt.Sprintf(
			"Invalid value '%s' for field '%s' with validation key '%s'.",
			err.FieldValue,
			err.FieldName,
			err.ValidationTag,
		)
		errStr = fmt.Sprintf("%s %s", errStr, fErrStr)
	}

	return errStr
}
