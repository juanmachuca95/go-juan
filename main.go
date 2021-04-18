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
	//Loading environment's variables
	loadEnv()

	port := os.Getenv("HEROKU_PORT")

	fmt.Printf("The host: %s", port)

	// Create a Fiber app
	app := fiber.New()
	
	/* App's Routes */
	api.SetupAppRoutes(app)

	/* Listen to port */
	if port == "" {
		_ = app.Listen(":3000")
	}else{
		_ = app.Listen("test")
	}
	
}

func loadEnv(){
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Cannot loading environment's variables")
	}
}