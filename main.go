package main

import (
	"github.com/gofiber/fiber/v2"
	"routers/router"
)

func main() {
	app := fiber.New()

	app.use(router.SetupRouter())
	app.Listen(":3000")
}