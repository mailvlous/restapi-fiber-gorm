package route

import (
	"restapi-fiber-gorm/handler"
	"restapi-fiber-gorm/config"
	"restapi-fiber-gorm/middleware"
	

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	
	r.Get("/", handler.HelloWorld)

	r.Post("/login", handler.Login)	

	r.Get("/getAllUser", middleware.Auth, handler.GetAllUser)
	r.Get("/getUser/:id", handler.GetUserById)
	r.Put ("/updateUser/:id", handler.UserUpdateName)
	r.Delete("/deleteUser/:id", handler.UserDelete)

	r.Static("/public", config.ProjectRootPath + "/public/assets")

	r.Post("/createUser", handler.CreateUser)
}
