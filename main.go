package main

import (
	"fmt"
	"os"
	"log"
	"go-juan/api"
	
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2"
)

// Primera api con Golang
func main() {
	loadEnv()

	env := os.Getenv("HOST")

	fmt.Printf("The host: %s", env)

	// Create a Fiber app
	app := fiber.New()
	
	/* App's Routes */
	api.SetupAppRoutes(app)

	/* Listen to port */
	_ = app.Listen(":3000")
}

func loadEnv(){
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Cannot loading environment's variables")
	}
}