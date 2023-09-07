package database


import (
	"database/sql"
	"fmt"
	// "go-litres/models"
	"go-litres/tools"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	Connection   *sql.DB
	Name string
	Host string
	Port string
	User string
	Pass string
}

func (d *DB) Connect() (connection *sql.DB) {
	connection, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", d.User, d.Pass, d.Host, d.Port, d.Name))
	tools.CheckErr(err)
	return
}

func (d *DB) Close() {
	d.Connection.Close()
}
