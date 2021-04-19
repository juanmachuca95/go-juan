package api 

import (
	D "go-juan/internal/database"
	"github.com/gofiber/fiber/v2"
	"database/sql"
	"log"
)

// Database instance
var db *sql.DB 

func getPadron(c *fiber.Ctx) error {
	db = D.Connect( db )

	// Get Employee list from database
	rows, err := db.Query( SelectPadron() )
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

	//Close db
	//defer db.Close()

	// Return Employees in JSON format
	return c.JSON(result)
}


func storePadron(c *fiber.Ctx) error {
	db = D.Connect( db )
	
	//New Padron struct
	u := new(Padron)

	// Parse body into struct
	if err := c.BodyParser(u); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Insert vote into database
	res, err := db.Query( InsertPadron() , u.Dni, u.Nombre, u.Apellido, u.Voto )
	if err != nil {
		return err
	}

	// Print result
	log.Println(res)

	//Close db
	defer db.Close()

	// Return vote in JSON format
	return c.JSON(u)
}