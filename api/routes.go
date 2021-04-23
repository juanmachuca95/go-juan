package api

import "github.com/gofiber/fiber/v2"


/*Mis routes*/
func SetupAppRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
<<<<<<< HEAD
        return c.SendString("Hola 👋! Bienvenidos!")
=======
        return c.SendString("Hola 👋! Bienvenido a mi API Golang. Juan Machuca")
>>>>>>> fed30be7ba28a0f1670967f3b54553a567ddd74b
    })

	app.Post("/login", login)

	app.Use(authRequired()).Get("/padron", getPadron) // READ PADRON

	app.Use(authRequired()).Post("/storepadron", storePadron) // CREATE PADRON

	app.Use(authRequired()).Post("/updatepadron", updatePadron) // UPDATE PADRON



}


