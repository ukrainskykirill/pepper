package types

import (
	"fmt"

	"github.com/dongri/phonenumber"
	"github.com/go-playground/validator/v10"
)

func NewValidator() *validator.Validate {
	validatorInstance := validator.New()
	err := validatorInstance.RegisterValidation("phone", validatePhone)
	return validatorInstance
}


func validatePhone(fl validator.FieldLevel) bool {
	land := phonenumber.GetISO3166ByNumber(fl.Field().String(), false)
	if land.Alpha2 != "RU" {
		return false
	} else { return true }
}
