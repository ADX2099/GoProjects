package val

import (
	"fmt"
	"lab01/api/classes"

	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

//CREACION DE INSTANCIA
func IniValidation() {
	validate = validator.New()
}
func ValidationLogin(user classes.UserLogin) bool {
	err_struct := validate.Struct(user)
	if err_struct != nil {
		for _, err := range err_struct.(validator.ValidationErrors) {

			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}
		return false
	} else {
		return true
	}
}
