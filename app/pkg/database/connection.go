package database

import (
	"fmt"

	"github.com/murbagus/hexapb-go/pkg/log"

	_ "github.com/lib/pq"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/rotisserie/eris"
)

// ConnectionDetail adalah konfigurasi struct untuk membangun suatu koneksi database
type ConnectionDetail struct {
	ConnectionName string
	Driver         DBDriver
	Username       string
	Host           string
	Port           string
	DBName         string
	Passwrod       string
}

// ConnectionConfig berisi konfigurasi konesi database
type ConnectionConfig struct {
	DefaultConnectionName string
	Connections           []ConnectionDetail
}

// DBConnection berisi data untuk melakuakan konesi database
type DBConnection struct {
	Driver       DBDriver
	DBName       string
	Conn         *sqlx.DB
	QueryBuilder squirrel.StatementBuilderType
}

// Connections merupakan objek berisi kumpulan koneksi yang di buat menggunakan factory function BuildConnections
type Connections struct {
	DefaultConnectionName string
	dbConnectionMap       map[string]*DBConnection
}

// BuildConnections merupakan factory function untuk struct Connection
func BuildConnections(cc *ConnectionConfig) (*Connections, error) {
	tmp := map[string]*DBConnection{}

	// Iterasi konesi
	// dan disimpan kedalam dbConnectionMap
	for _, v := range cc.Connections {
		log.ConsoleInfo("membangun koneksi database", v.ConnectionName)

		dd := v.Driver.Detail(v.Username, v.Passwrod, v.Host, v.Port, v.DBName)

		conn, err := sqlx.Connect(v.Driver.String(), dd.DSN)
		if err != nil {
			// return nil, eris.New("Terjadi kelsalahan saat membangun koneksi")
			return nil, err
		}

		tmp[v.ConnectionName] = &DBConnection{
			Driver:       v.Driver,
			DBName:       v.DBName,
			Conn:         conn,
			QueryBuilder: squirrel.StatementBuilder.PlaceholderFormat(dd.PlacehoderFormat),
		}
	}

	return &Connections{
		DefaultConnectionName: cc.DefaultConnectionName,
		dbConnectionMap:       tmp,
	}, nil
}

func (c *Connections) connect(name string) (*DBConnection, error) {
	if val, ok := c.dbConnectionMap[name]; ok {
		return val, nil
	}

	return nil, eris.New(fmt.Sprint("Koneksi database ", name, " tidak ditemukan"))
}

// Default mengembalikan DBConnection default
func (c *Connections) GetConnectionDefault() (*DBConnection, error) {
	conn, err := c.connect(c.DefaultConnectionName)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// Get mengembalikan connectionData sesuai dengan yang diminta
func (c *Connections) GetConnection(name string) (*DBConnection, error) {
	conn, err := c.connect(name)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
