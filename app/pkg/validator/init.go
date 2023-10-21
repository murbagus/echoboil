package validator

import (
	"reflect"

	"github.com/go-playground/locales/id_ID"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	t_id "github.com/murbagus/hexapb-go/pkg/validator/translation/id"
)

var (
	vali  *validator.Validate
	trans ut.Translator
)

func init() {
	vali = validator.New()

	id := id_ID.New()
	universalTranslator := ut.New(id, id)

	trans, _ = universalTranslator.GetTranslator("id_ID")

	t_id.RegisterDefaultTranslations(vali, trans)

	// Tag ini nanti akan digunakan
	// untuk menjadi nama alternatif dalam pesan error
	vali.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("govaf")
	})
}
