package utils

import (
	"net/http"
)

const (
	HTTP_RESPONSE_CODE_SUCCESS          = 0
	HTTP_RESPONSE_CODE_ERROR_MESSAGE    = 1 // HTTP_ERROR_RESPONS_CODE_MESSAGE untuk http respons yang hanya memberikan pesan error
	HTTP_RESPONSE_CODE_ERROR_VALIDATION = 2 // HTTP_ERROR_RESPONS_CODE_VALIDATION untuk http respons data gagal validasi
)

// HttpResponseTemplate templat respons khusus untuk http
type HttpResponseTemplate struct {
	Code    int         `json:"kode,omitempty"`
	Message string      `json:"pesan,omitempty"`
	Details interface{} `json:"detail,omitempty"`
}

// HttpErrorResponseBinding mengembalikan http responss Code dan pesan error
func HttpErrorResponseBinding(err error) (int, HttpResponseTemplate) {
	return http.StatusUnprocessableEntity, HttpResponseTemplate{
		Code:    HTTP_RESPONSE_CODE_ERROR_MESSAGE,
		Message: "Tipe data tidak sesuai",
		Details: "Hubungi developer karena terlah terjadi kesalahan pada sistem",
	}
}
