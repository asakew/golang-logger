// https://docs.gofiber.io/api/middleware/logger/
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Create a new Fiber app
	app := fiber.New()

	// Set up the logging middleware
	//app.Use(logger.New()) // default

	// Custom logger
	app.Use(logger.New(logger.Config{
		Format:     "${time} - ${status} - ${method} - ${ip} - ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
	}))

	// Define a simple route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! 🌍")
	})

	// Start the server
	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
