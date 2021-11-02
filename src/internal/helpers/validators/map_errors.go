package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func MapErrorsToMessages(err error) (errors []string) {
	for _, err := range err.(validator.ValidationErrors) {
		errors = append(
			errors,
			fmt.Sprintf("Field %v failed validation for %v", err.Field(), err.Tag()))
	}
	return
}
