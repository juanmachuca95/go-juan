package api

import "github.com/gofiber/fiber/v2"


/*Mis routes*/
func SetupAppRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hola 👋! Bienvenidos!")
    })

	app.Post("/login", login)

	app.Use(authRequired()).Get("/padron", getPadron) // READ PADRON

	app.Use(authRequired()).Post("/storepadron", storePadron) // CREATE PADRON

	app.Use(authRequired()).Post("/updatepadron", updatePadron) // UPDATE PADRON



}


