package handler

import (
	"github.com/gofiber/fiber/v2"
	"realApi/app/api"
)

func ApiRoutes() {
	app := fiber.New()
	app.Get("/", api.Welcome)
	app.Get("/customers", api.GetAllCustomer)
	app.Get("/customer/:email", api.GetByUsernameWith)
	app.Post("/customer", api.CreateCustomer)
	app.Post("/customers/update", api.UpdateCustomer)
	app.Delete("/customer/:username", api.DeleteCustomer)

	err := app.Listen(":4000")
	if err != nil {
		return
	}
}
