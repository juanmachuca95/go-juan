package api

import "github.com/gofiber/fiber/v2"
import _ "github.com/go-sql-driver/mysql"
import _"log"
import _"fmt"
import _"database/sql"

/*Mis routes*/
func SetupAppRoutes(app *fiber.App) {

	// Get all records from MySQL
	app.Get("/padron", getPadron)
	// Add record into MySQL
	app.Post("/storepadron", storePadron)

}