package main

import (
	"github.com/gofiber/fiber/v2"
	"learngo/routers"
	"learngo/db"
	"fmt"
)

func main() {
	fromDb:= database.Init()
	fmt.Println(fromDb,"frm db")
	app := fiber.New()

	router.SetupRouter(app)
	app.Listen(":3000")
}