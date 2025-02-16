package middleware

import (
	"github.com/gofiber/fiber/v2"
	"restapi-fiber-gorm/utils"
	"fmt"
)

func Auth(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")

	fmt.Println("Received x-token header:", token) // Debugging

	if token == "" {
		fmt.Println("Token is missing!") // Debugging
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated1",
		})
	}

	claims, err := utils.DecodeToken(token)
	if err != nil {
		fmt.Println("Token verification failed:", err) // Debugging
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated2",
		})
	}

	role, ok := claims["role"].(string)
	if !ok {
		fmt.Println("Role field is missing or invalid") // Debugging
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated3",
		})
	}

	if role != "admin" {
		fmt.Println("Unauthorized role:", role) // Debugging
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated4",
		})
	}

	return ctx.Next()
}


