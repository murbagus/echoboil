package validator

type ErrorsJson map[string]interface{}
type FieldRules map[string]interface{}
type VarArrMap []map[string]interface{}
type ValidateArrayFunc func(gv GoVa)

type GoVa interface {
	GetFieldErrors() []fieldError

	AppendFieldErrors(field string, message string)
}

type fieldError struct {
	field   string
	message string
}

type gova struct {
	fes []fieldError
}

func New() *gova {
	return &gova{
		fes: []fieldError{},
	}
}

// HasErrors mengembalikan true jika terdapat field error
func (gv *gova) HasErrors() bool {
	return len(gv.fes) > 0
}

// GetFieldErrors digunakan untuk mendapatkan field error
func (gv *gova) GetFieldErrors() []fieldError {
	return gv.fes
}

// GetFieldErrors digunakan untuk mendapatkan field error
// yang kompatibel dengan respons json framework echo
func (gv *gova) ResponsErrorsJson() ErrorsJson {
	tmp := make(ErrorsJson)
	for _, fe := range gv.GetFieldErrors() {
		tmp[fe.field] = fe.message
	}

	return tmp
}

// AppendFieldErrors digunakan untuk menambah field error
func (gv *gova) AppendFieldErrors(field string, message string) {
	gv.fes = append(gv.fes, fieldError{
		field:   field,
		message: message,
	})
}
