package api

import (

	"github.com/gofiber/fiber/v2"
	_ "github.com/dgrijalva/jwt-go"
	jwtware "github.com/gofiber/jwt/v2"
	"os"

)

func authRequired() fiber.Handler {
	
	var jwtSecret string = os.Getenv("TOKEN_KEY")
	print(jwtSecret)
	return jwtware.New(jwtware.Config{	
		SigningKey: []byte(jwtSecret),
	})
	
}