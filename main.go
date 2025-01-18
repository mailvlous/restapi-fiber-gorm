package main

import (
	"github.com/gofiber/fiber/v2"
	"restapi-fiber-gorm/route"
	"restapi-fiber-gorm/database"
	"restapi-fiber-gorm/database/migration"
	
)

func main() {
	database.DatabaseInit()
	migration.RunMigration()

	app := fiber.New()

	route.RouteInit(app)

	app.Listen(":8080")
}