package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func Validate(gv GoVa, variable interface{}, name string, rules string) {
	verr := vali.Var(variable, rules)

	if verr != nil {
		ve := verr.(validator.ValidationErrors)

		gv.AppendFieldErrors(name, strings.TrimSpace(ve[0].Translate(trans)))
	}
}

func ValidateArrMap(gv GoVa, arrMap VarArrMap, name string, fr FieldRules) error {
	for mapIndex, mapp := range arrMap {

		for field, rules := range fr {

			if val, ok := mapp[field]; ok {
				switch reflect.TypeOf(rules).Name() {

				case "string":
					Validate(gv, val, fmt.Sprintf("%s.%d.%s", name, mapIndex, field), rules.(string))

				case "FieldRules":
					if t, ok := val.(VarArrMap); ok {

						err := ValidateArrMap(gv, t, fmt.Sprintf("%s.%d.%s", name, mapIndex, field), rules.(FieldRules))
						if err != nil {
							return err
						}

					} else {
						return fmt.Errorf("field %s.%s bukan bertipe VarArrMap", name, field)
					}

				default:
					return fmt.Errorf("tipe rules yang ditulis pada field %s.%s tidak sesuai", name, field)
				}
			}
		}
	}

	return nil
}
