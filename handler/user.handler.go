package handler

import (
	"github.com/gofiber/fiber/v2"
)

// CreateUserHandler handles the creation of a new user

func HelloHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}