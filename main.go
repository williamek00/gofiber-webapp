package main

import (
	"github.com/gofiber/fiber/v2"
	"learngo/routers"
	"learngo/db"
	"fmt"
)

func main() {
	fromDb,_:= database.Db()
	fmt.Println(fromDb,"main.go")
	app := fiber.New()
	router.SetupRouter(app)
	app.Listen(":3000")
}