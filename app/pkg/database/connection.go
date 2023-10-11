package database

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/mitchellh/mapstructure"
	"github.com/murbagus/hexapb-go/pkg/dir"
	"github.com/murbagus/hexapb-go/pkg/log"
	"github.com/rotisserie/eris"
	"github.com/spf13/viper"
)

type connectionData struct {
	Driver DBDriver
	DBName string
	Conn   *sqlx.DB
}

type connectionConfig struct {
	Driver   string `mapstructure:"driver"`
	Username string `mapstructure:"username"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	DBName   string `mapstructure:"db_name"`
	Passwrod struct {
		Value      string `mapstructure:"value"`
		IsFilePath bool   `mapstructure:"is_file_path"`
	} `mapstructure:"password"`
}

var connection map[string]*connectionData = map[string]*connectionData{}

func init() {
	build()
}

func build() {
	viper.SetConfigName("database")
	viper.SetConfigType("toml")
	viper.AddConfigPath(dir.GetConfigPath())

	err := viper.ReadInConfig()
	if err != nil {
		err = eris.Wrap(err, "Gagal membaca config database")
		log.ConsoleFatal(err)
	}

	var cc map[string]connectionConfig
	err = mapstructure.Decode(viper.Get("connection"), &cc)
	if err != nil {
		err = eris.Wrap(err, "Gagal menerjemahkan config database")
		log.ConsoleFatal(err)
	}

	// Melakukan iterasi koneksi yang ada
	// di dalam file konfigurasi.
	//
	// Dalam tahap ini juga ditentukan value dari
	// konfigrasi password database, apakah itu adalah
	// password atau merupakan path file yang berisi
	// password.
	for k, v := range cc {
		log.ConsoleInfo("Membuat koneksi database", k, "...")

		driver, err := GetDriver(v.Driver)
		if err != nil {
			log.ConsoleFatal(err)
		}

		password := v.Passwrod.Value
		if v.Passwrod.IsFilePath {
			f, err := os.ReadFile(v.Passwrod.Value)

			if err != nil {
				err = eris.New(fmt.Sprint("Gagal membaca path file password koneksi database", k))
				log.ConsoleFatal(err)
			}

			password = string(f)
		}

		connection[k] = &connectionData{
			Driver: driver,
			DBName: v.DBName,
			Conn:   create(driver, v.Username, password, v.Host, v.Port, v.DBName),
		}
	}
}

func create(driver DBDriver, username string, password string, host string, port string, dbName string) *sqlx.DB {
	conn, err := sqlx.Connect(driver.String(), driver.DSN(username, password, host, port, dbName))

	if err != nil {
		err = eris.Wrap(err, "Terjadi kelsalahan saat membuat koneksi")
		log.ConsoleFatal(err)
	}

	return conn
}

func connect(name string) (*connectionData, error) {
	if val, ok := connection[name]; ok {
		return val, nil
	}

	return nil, eris.New(fmt.Sprint("Koneksi database ", name, " tidak ditemukan"))
}

// Default mengembalikan connectionData default
func Default() *connectionData {
	res, err := connect("default")
	if err != nil {
		log.ConsoleFatal(err)
	}

	return res
}

// Get mengembalikan connectionData sesuai dengan yang diminta
func Get(name string) *connectionData {
	res, err := connect(name)
	if err != nil {
		log.ConsoleFatal(err)
	}

	return res
}
