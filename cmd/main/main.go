package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

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

func main() {
	// Connect with database
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	// Create a Fiber app
	app := fiber.New()

	// Get all records from MySQL
	app.Get("/padron", func(c *fiber.Ctx) error {
		// Get Employee list from database
		rows, err := db.Query("SELECT dni, nombre, apellido, voto FROM padron")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		defer rows.Close()
		result := Padrones{}

		for rows.Next() {
			padron := Padron{}
			if err := rows.Scan(&padron.Dni, &padron.Nombre, &padron.Apellido, &padron.Voto); err != nil {
				return err // Exit if we get an error
			}

			// Append Employee to Employees
			result.Padrones = append(result.Padrones, padron)
		}
		// Return Employees in JSON format
		return c.JSON(result)
	})

	// Add record into MySQL
	app.Post("/storepadron", func(c *fiber.Ctx) error {
		//New Padron struct
		u := new(Padron)

		// Parse body into struct
		if err := c.BodyParser(u); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		// Insert vote into database
		res, err := db.Query("INSERT INTO padron (dni, nombre, apellido, voto ) VALUES (?, ?, ?, ?)", u.Dni, u.Nombre, u.Apellido, u.Voto )
		if err != nil {
			return err
		}

		// Print result
		log.Println(res)

		// Return vote in JSON format
		return c.JSON(u)
	})

	
	log.Fatal(app.Listen(":3000"))
}