package database

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import _ "go-juan/internal/logs"


type MySqlClient struct {
	*sql.DB
}

func NewMySQLClient() *MySqlClient {

	/*Conección a la base de datos*/
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gopruebas")

	//Manejo de error
	if err != nil {		
		panic(err.Error())
	}

	defer db.Close()
	
	fmt.Println("Succesfully conected to MySQL database")
	
	return &MySqlClient{db}
}	