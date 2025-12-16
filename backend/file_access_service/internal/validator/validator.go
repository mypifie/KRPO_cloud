package validator

import (
	"strings"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func NewValidator() *CustomValidator {
	v := validator.New()
	return &CustomValidator{Validator: v}

}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func GetValidationErrors(err error) error {
	valErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}
	var msgs []string
	for _, fieldErr := range valErrors {
		msg := fmt.Sprintf("%s: %s",
			strings.ToLower(fieldErr.Field()),
			getErrorText(fieldErr),
		)
		msgs = append(msgs, msg)
	}

	return fmt.Errorf("validation failed: %s", strings.Join(msgs, ", "))
}

func getErrorText(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "this field is required"
	default:
		return "invalid"
	}
}
