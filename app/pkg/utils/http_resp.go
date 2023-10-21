package utils

import "net/http"

// HttpResponsErrorBinding mengembalikan http respons code dan pesan error
func HttpResponsErrorBinding(err error) (int, map[string]string) {
	return http.StatusBadRequest, map[string]string{
		"message": "Tipe data payload tidak sesuai",
	}
}
