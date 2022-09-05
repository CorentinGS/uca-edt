package main

import (
	"fmt"
	"github.com/corentings/uca-edt/pkg/parsing"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Printf("Hello, world.")
	fmt.Printf("This is our special EDT project for UCA.\n")

	parsing.Parse()

	// Create the app
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
