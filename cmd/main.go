package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Perfect Panel",
	})

	app.Route("/api", func(api fiber.Router) {
		//api v1
		v1 := api.Group("/v1")
		v1.Route("/auth", func(auth fiber.Router) {

		})
	})

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
