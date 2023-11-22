package validator

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// ValidateStruct digunakan untuk melakukan validasi struct dengan
// rule yang telah ditentukan dalam modul validatro v10
func ValidateStruct(gv *gova, s interface{}) {
	verr := vali.Struct(s)

	if verr != nil {
		ve := verr.(validator.ValidationErrors)

		for _, fe := range ve {
			fieldName := fe.Namespace()
			fieldName = strings.ReplaceAll(fieldName, "[", ".")
			fieldName = strings.ReplaceAll(fieldName, "]", "")

			gv.AppendFieldErrors(fieldName, fe.Translate(trans))
		}
	}
}
