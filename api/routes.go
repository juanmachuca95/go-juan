package api

import "github.com/gofiber/fiber/v2"


/*Mis routes*/
func SetupAppRoutes(app *fiber.App) {

	app.Post("/login", login)
	// Get all records from MySQL grp.Use(jwtMiddleware(tokenKey)).Post("/permisos", s.PermisoHandler)
	app.Use(authRequired()).Get("/padron", getPadron)
	// Add record into MySQL
	app.Use(authRequired()).Post("/storepadron", storePadron)

	//app.Get("")
	app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hola 👋! Bienvenido a mi API Golang")
    })

}


