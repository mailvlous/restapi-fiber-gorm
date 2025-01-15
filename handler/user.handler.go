package handler

import (
	"github.com/gofiber/fiber/v2"

	"restapi-fiber-gorm/model/entity"
	"restapi-fiber-gorm/database"
	"restapi-fiber-gorm/model/request"

	"github.com/go-playground/validator/v10"
)

func HelloWorld(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, World!")
}

// CreateUserHandler handles the creation of a new user

func GetAllUser(ctx *fiber.Ctx) error {
	var users []entity.Users

    result := database.DB.Find(&users)
    if result.Error != nil {
        return ctx.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
    }

    return ctx.JSON(users)
}

func CreateUser(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	// Request Validation
	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": "Validation error",
			"error": errValidate.Error(),
		})
	}

	newUser := entity.Users{
		Name: user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}

	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser!= nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Could not create user", 
			"data": nil,
		})
	}

	return ctx.JSON(fiber.Map{
		"status": "success",
		"message": "User successfully created",
		"data": newUser,
	})

}

func GetDataById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user entity.Users

	result := database.DB.First(&user, id)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	return ctx.JSON(user)
}