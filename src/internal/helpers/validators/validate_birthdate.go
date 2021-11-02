package validators

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/ppcamp/go-graphql-with-auth/internal/utils"
)

var birthdateValidation validator.Func = func(fl validator.FieldLevel) bool {
	birthdate, ok := fl.Field().Interface().(time.Time)
	if ok {
		return utils.IsAValidBirthDate(birthdate)
	}
	return false
}
