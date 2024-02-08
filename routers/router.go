package router

import "github.com/gofiber/fiber/v2"

func SetupRouter(app *fiber.App) {
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })
    // Add more routes and middleware as needed
}
