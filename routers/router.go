package router

import (
	userhandler "learngo/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	app.Get("/", userhandler.GetUsers)
	app.Get("/user/:id", userhandler.GetUser)
	app.Post("/user", userhandler.CreateUser)
	app.Delete("user/:id", userhandler.DeleteUser)
}
