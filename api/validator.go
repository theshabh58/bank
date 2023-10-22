package api

import (
	"bank/util"

	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldlevel validator.FieldLevel) bool {
	if currency, ok := fieldlevel.Field().Interface().(string); ok {
		// check if currency is supported
		return util.IsSupportedCurrency(currency)
	}
	return false
}
