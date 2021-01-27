package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin/binding"
)

var validType validator.Func = func(fl validator.FieldLevel) bool {
	validTypes := []string{"credit", "debit"}

	for _, validType := range validTypes {
		if validType == fl.Field().String() {
			return true
		}
	}

	return false
}

var amount validator.Func = func(fl validator.FieldLevel) bool {
	if fl.Field().Float() > 0 {
		return true
	}

	return false
}

func Register() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("validType", validType)
		v.RegisterValidation("amount", amount)
	}
}

