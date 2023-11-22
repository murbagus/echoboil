package database

import (
	"fmt"

	"github.com/Masterminds/squirrel"
)

type driverDetail struct {
	DSN              string
	PlacehoderFormat squirrel.PlaceholderFormat
}

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

func (d DBDriver) Detail(username string, password string, host string, port string, dbName string) *driverDetail {
	switch d {
	case Postgres:
		return &driverDetail{
			DSN:              fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, dbName),
			PlacehoderFormat: squirrel.Dollar,
		}
	default:
		return &driverDetail{
			DSN:              fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, dbName),
			PlacehoderFormat: squirrel.Dollar,
		}
	}
}
