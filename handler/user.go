package userhandler

import (
    "github.com/gofiber/fiber/v2"
    "learngo/model"
)

func GetUser(c *fiber.Ctx) error {
    users, err := usermodel.GetUser()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch users")
    }

    return c.JSON(users)
}