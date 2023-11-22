package utils

// Must adalah fungsi untuk mengambil nilai balikan pertama pada sebuah fungsi
// yang mengembalikan dua nilai (nilai error di akhir balikan) tanpa balikan nilai error
func Must[T any](res T, err error) T {
	if err != nil {
		panic(err)
	}

	return res
}
