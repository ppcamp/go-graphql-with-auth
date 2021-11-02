package validators

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func init() {
	setupValidator()
}

func setupValidator() {
	if Validator == nil {
		Validator = validator.New()
		Validator.SetTagName("binding")
		Validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		Validator.RegisterValidation("birthdate", birthdateValidation)
	}
}
