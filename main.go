package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/riyan-eng/ngitung-dhuit/config"
	"github.com/riyan-eng/ngitung-dhuit/module/finance"
)

func init() {
	os.Setenv("TZ", "Asia/Jakarta")
	config.Environment()
	config.DatabaseConnection()
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("help")
	})

	finance.Setup(app)
	app.Listen(os.Getenv("SERVER_URL"))
}
