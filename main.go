package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamek00/gofiber-webapp/routers/router"
)

func main() {
	app := fiber.New()

	router.SetupRouter(app)
	app.Listen(":3000")
}