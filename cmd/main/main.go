package main

import (
	"go-juan/api"
	
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

// Primera api con Golang
func main() {
	// Create a Fiber app
	app := fiber.New()
	
	/* App's Routes */
	api.SetupAppRoutes(app)

	/* Listen to port */
	_ = app.Listen(":3000")
}