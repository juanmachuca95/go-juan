package api

import "github.com/gofiber/fiber/v2"


/*Mis routes de la api*/
func SetupAppRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hola 👋! Bienvenido a mi API Golang. Version 0.1")
    })

	app.Post("/login", login)

	app.Use(authRequired()).Get("/padron", getPadron) // READ PADRON

	app.Use(authRequired()).Post("/storednipadron", storeDniPadron) // CREATE PADRON POR DNI NO ENCONTRADO

	//app.Use(authRequired()).Post("/storepadron", storePadron) // CREATE PADRON


	app.Use(authRequired()).Post("/updatepadron", updatePadron) // UPDATE PADRON



}


