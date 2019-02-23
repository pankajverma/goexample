package customevalidator

import (
	"reflect"

	"gopkg.in/go-playground/validator.v8"
)

/*
PhoneNumerValidator ...
*/
func PhoneNumerValidator(v *validator.Validate,
	topStruct reflect.Value,
	currentStructOrField reflect.Value,
	field reflect.Value,
	fieldType reflect.Type,
	fieldKind reflect.Kind,
	param string) bool {

	if number, ok := field.Interface().(string); ok {
		if len(number) >= 10 && len(number) <= 12 {
			return true
		}
	}
	return false
}

/*
RegisterCustomeValidators ...
*/
func RegisterCustomeValidators(v *validator.Validate) {
	// NOTES: using the same tag name as an existing function
	//        will overwrite the existing one
	v.RegisterValidation("PhoneNumerValidator", PhoneNumerValidator)

}
