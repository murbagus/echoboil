package validator

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

// Err digunakan untuk mendapatkan field error
func (gv *gova) Err() []fieldError {
	return gv.fes
}

// ErrJSON digunakan untuk mendapatkan field error
// yang kompatibel dengan respons json framework echo
func (gv *gova) ErrJSON() map[string]interface{} {
	tmp := make(map[string]interface{})
	for _, fe := range gv.Err() {
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
