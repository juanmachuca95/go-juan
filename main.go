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

	// Create a Fiber app
	app := fiber.New()
	
	/* App's Routes */
	api.SetupAppRoutes(app)

	/* Listen to port */
	port := os.Getenv("HEROKU")
    if port == "" {
        port = "3000"
    } else {
		port = os.Getenv("HEROKU")
	}
	
	fmt.Printf("The host: %s", port)

    _ = app.Listen(":"+port)
		
}

func loadEnv(){
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Cannot loading environment's variables")
	}
}