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

	if os.Getenv("DEFAULT_CONNECTION") == true {
		fmt.Println("Utilizar la connection en local")
		// Database settings
		var (
			host     = os.Getenv("HOST_LOCAL")
			user     = os.Getenv("USER_LOCAL")
			password = os.Getenv("PASSWORD_LOCAL")
			dbname   = os.Getenv("DATABASE_LOCAL")
		)
	}else {
		// Database settings
		var (
			host     = os.Getenv("HOST")
			user     = os.Getenv("USER")
			password = os.Getenv("PASSWORD")
			dbname   = os.Getenv("DATABASE")
		)
		fmt.Println("Utilizar la conección de producción")
	}
	

	fmt.Printf("%s", host)
	fmt.Println("***********")
	fmt.Printf("%s", user)
	fmt.Println("***********")
	fmt.Printf("%s", password)
	fmt.Println("***********")
	fmt.Printf("%s", dbname)

	var err error
	// Use DSN string to open
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, dbname))
	if err != nil {
		panic(err)
	}

	return db
}