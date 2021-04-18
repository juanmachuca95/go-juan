package database

import "fmt"
import "os"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import _ "go-juan/internal/logs"

// Database instance
var db *sql.DB


// Connect function
func Connect(db *sql.DB) *sql.DB {

	// Database settings
	var (
		host     = os.Getenv("HOST")
		user     = os.Getenv("USER")
		password = os.Getenv("PASSWORD")
		dbname   = os.Getenv("DATABASE")
	)

	fmt.Printf("%s", host)

	var err error
	// Use DSN string to open
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, dbname))
	if err != nil {
		panic(err)
	}

	return db
}