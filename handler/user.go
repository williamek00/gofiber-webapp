package userhandler

import (
	usermodel "learngo/model"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	users, err := usermodel.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch users")
	}
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	user := usermodel.GetUser(c)
	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	user := usermodel.CreateUser(c)
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	result := usermodel.DeleteUser(c)
	return c.JSON(result)
}
