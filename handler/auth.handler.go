package handler

import (
	"log"
	"restapi-fiber-gorm/database"
	"restapi-fiber-gorm/model/entity"
	"restapi-fiber-gorm/model/request"
	"restapi-fiber-gorm/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

)

func Login(ctx *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)

	if err := ctx.BodyParser(loginRequest); err != nil {
		return err
	}
	log.Println(loginRequest)

		// VALIDASI REQUEST
	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	//	CHECK AVAILABLE USER\
	var user entity.Users
	err := database.DB.First(&user, "email = ?", loginRequest.Email).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found",
			"data":    nil,
		})
	}


	//	CHECK PASSWORD
	isValid := utils.CheckPasswordHash(loginRequest.Password, user.Password)

	
	if !isValid {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found",
			"data":    nil,
		})
	}

	return ctx.JSON(fiber.Map{
		"token": "secret",
	})

}