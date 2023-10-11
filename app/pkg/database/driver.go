package database

import (
	"fmt"

	"github.com/rotisserie/eris"
)

// DBDriver adalah enum untuk driver yang dapat dipilih
// pada saat membuat koneksi database baru
type DBDriver int8

const (
	Undefined DBDriver = iota
	Postgres
)

func (d DBDriver) String() string {
	switch d {
	case Postgres:
		return "postgres"
	default:
		return "postgres"
	}
}

// DSN mengembalikan string DSN stiap driver
func (d DBDriver) DSN(username string, password string, host string, port string, dbName string) string {
	switch d {
	case Postgres:
		return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, dbName)
	default:
		return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, dbName)
	}
}

// GetDriver memilih driver dari bentuk string
func GetDriver(driver string) (DBDriver, error) {
	switch driver {
	case "postgres":
		return Postgres, nil
	default:
		return Undefined, eris.New(fmt.Sprint("Driver database ", driver, " tidak ditemukan"))
	}
}
