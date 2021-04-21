package api

import (

	"github.com/gofiber/fiber/v2"
	_ "github.com/dgrijalva/jwt-go"
	jwtware "github.com/gofiber/jwt/v2"
	"os"

)

// Autorización para realizar acciones en el servicio
/* 
	Filter: nil,
	SuccessHandler: nil,
	SigningKeys: nil,
	SigningMethod: "",
	ContextKey: "",
	Claims: nil,
	TokenLookup: "",
	AuthScheme: "", 
	*/
func authRequired() fiber.Handler {
	
	var jwtSecret string = os.Getenv("TOKEN_KEY")
	print(jwtSecret)
	return jwtware.New(jwtware.Config{	
		/* ErrorHandler: func(ctx *fiber.Ctx, err error){
			ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":"Unauthorized",
			})
		}, */
		SigningKey: []byte(jwtSecret),
	})
}