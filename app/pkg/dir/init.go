package dir

import (
	"path"
	"runtime"
)

var (
	curpath  string
	basepath string
	cfgpath  string
	logpath  string
)

func init() {
	_, filename, _, _ := runtime.Caller(0)

	curpath = path.Dir(filename)
	basepath = path.Join(curpath, "../../")

	cfgpath = path.Join(basepath, "/cfg")
	logpath = path.Join(basepath, "/log")
}

// GetBasePath mengembalikan string path direktori proyek
func GetBasePath() string {
	return basepath
}

// GetConfigPath mengembalikan string path direktori konfigurasi
func GetConfigPath() string {
	return cfgpath
}

// GetLogPath mengembalikan string path direktori berkas log
func GetLogPath() string {
	return logpath
}
