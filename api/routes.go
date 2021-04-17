package api

import "github.com/gofiber/fiber/v2"
import _ "github.com/go-sql-driver/mysql"
import "log"
import "fmt"
import "database/sql"

// Database instance
var db *sql.DB

// Database settings
const (
	host     = "localhost"
	port     = 3306 // Default port
	user     = "root"
	password = ""
	dbname   = "vi"
)

// Employee struct
type Padron struct {
	Dni         int     `json:"dni"`
	Nombre      string  `json:"nombre"`
  	Apellido    string  `json:"apellido"`
  	Voto        bool    `json:"voto"`
}

// Employees struct
type Padrones struct {
	Padrones []Padron `json:"votantes"`
}

// Connect function
func Connect() error {
	var err error
	// Use DSN string to open
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, dbname))
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}


/*Mis routes*/
func SetupAppRoutes(app *fiber.App) {

	// Connect with database
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	// Get all records from MySQL
	app.Get("/padron", getPadron)
	// Add record into MySQL
	app.Post("/storepadron", storePadron)

}