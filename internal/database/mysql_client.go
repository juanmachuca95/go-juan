package database

import "database/sql"
import "github.com/go-sql-driver/mysql"
import "go-juan/internal/logs"


type MySqlClient struct {
	*sql.DB
}

func NewMySQLClient() *MySqlClient {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gopruebas")

	if err != nil {
		logs.Error("cannot create mysql client")
		panic(err)
	}

	err = db.Ping()

	if err != nil {

	}

	return &MySqlClient{db}
}	