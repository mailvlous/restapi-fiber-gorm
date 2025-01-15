package route

import (
	"restapi-fiber-gorm/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/", handler.HelloWorld)
	r.Get("/getAllUser", handler.GetAllUser)

	r.Post("/createUser", handler.CreateUser)
}
