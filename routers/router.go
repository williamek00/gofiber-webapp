package router

import (
    "github.com/gofiber/fiber/v2"
    "learngo/handler"
)

func SetupRouter(app *fiber.App) {
    app.Get("/",userhandler.GetUser)
}
