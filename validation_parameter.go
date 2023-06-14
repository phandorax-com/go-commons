package resttemplate

import (
	"github.com/go-playground/validator/v10"
	"github.com/subchen/go-log"
)

var validate *validator.Validate

func validationParameter(parameter, evaluation string) bool {
	if errors := validate.Var(parameter, evaluation); errors != nil {
		log.Errorf("Parameter: %s - Error: %s", parameter, errors)
		return false
	}
	return true
}
