package validator

import (
	"regexp"

	validator "gopkg.in/go-playground/validator.v9"
)

// CustomValidator needed to use the go-playground/validator pkg
type CustomValidator struct {
	validator *validator.Validate
}

// Validate small function needed to use go-playground/validator
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// CreateValidator creates the validator to set in the echo server struct
func CreateValidator() *CustomValidator {
	v := validator.New()
	v.RegisterValidation("validchannelname", isValidChannelName)
	return &CustomValidator{validator: v}
}

// IsAlphanumericOrDash checks if a string is a-z, A-Z, 0-9 or a -
func IsAlphanumericOrDash(s string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9-]+$")

	return re.MatchString(s)
}

func isValidChannelName(fl validator.FieldLevel) bool {
	return IsAlphanumericOrDash(fl.Field().String())
}
