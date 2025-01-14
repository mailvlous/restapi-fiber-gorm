package route

import (
	"github.com/gofiber/fiber/v2"
	"restapi-fiber-gorm/handler"
)


func RouteInit(r *fiber.App) {
	r.Get("/", handler.HelloHandler)
}