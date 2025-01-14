package main

import (
	"github.com/gofiber/fiber/v2"
	"restapi-fiber-gorm/route"
)

func main() {
	app := fiber.New()

	route.RouteInit(app)

	app.Listen(":8080")
}