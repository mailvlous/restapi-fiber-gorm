package handler

import (
	"log"
	"restapi-fiber-gorm/database"
	"restapi-fiber-gorm/model/entity"
	"restapi-fiber-gorm/model/request"
	"restapi-fiber-gorm/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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

	// GENERATE JWT

	claims := jwt.MapClaims{}
	claims["id"] = user.Id
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	if user.Email == "mail@mail.com" {
		claims["role"] = "admin"
	} else {
		claims["role"] = "user"
	}

	token, errGenerateToken := utils.GenerateToken(&claims)

	if errGenerateToken != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed",
			"error":   errGenerateToken.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})

}