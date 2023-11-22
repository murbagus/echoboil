package utils

import (
	"net/http"
)

// E untuk membuat konstan error pada package port
// ini dikarenakan tidak dapat membuat error.New() pada konstan
// maka sengaja dibuat tipe E berdasar string dengan
// menerapkan interface error
type E string

func (e E) Error() string { return string(e) }

const (
	ERR_SERVICE_T_DEPENDENT_DATA_VALIDATION E = "Data yang dibutuhkan tidak valid"
	ERR_SERVICE_T_FLOW_VALIDATION           E = "Alur service berhenti karena validasi lanjutan"
	ERR_SERVICE_T_REPO                      E = "Terjadi kesalahan pada saat integrasi dengan repository"
	ERR_SERVICE_T_DATA_NOT_FOUND            E = "Data tidak ditemukan"
)

// ErrService adalah struct yang digunakan pada saat pengembalian error pada service
type ServiceErr struct {
	Type    error
	Message string
	Details interface{}
}

func (se ServiceErr) ResponseForHttp() (int, HttpResponseTemplate) {
	switch se.Type {
	case ERR_SERVICE_T_DEPENDENT_DATA_VALIDATION:
		return http.StatusUnprocessableEntity, HttpResponseTemplate{
			Code:    HTTP_RESPONSE_CODE_ERROR_VALIDATION,
			Message: se.Message,
			Details: se.Details,
		}
	case ERR_SERVICE_T_DATA_NOT_FOUND:
		return http.StatusNotFound, HttpResponseTemplate{
			Code:    HTTP_RESPONSE_CODE_ERROR_MESSAGE,
			Message: se.Message,
			Details: se.Details,
		}
	case ERR_SERVICE_T_REPO:
		return 500, HttpResponseTemplate{
			Code:    HTTP_RESPONSE_CODE_ERROR_MESSAGE,
			Message: se.Message,
			Details: se.Details,
		}
	default:
		return http.StatusBadRequest, HttpResponseTemplate{
			Code:    HTTP_RESPONSE_CODE_ERROR_MESSAGE,
			Message: se.Message,
			Details: se.Details,
		}
	}
}
