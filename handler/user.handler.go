package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"restapi-fiber-gorm/database"
	"restapi-fiber-gorm/model/entity"
	"restapi-fiber-gorm/model/request"
	"restapi-fiber-gorm/model/response"
	"restapi-fiber-gorm/utils"

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

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Could not create user",
			"data": nil,
		})
	}

	newUser.Password = hashedPassword

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

func GetUserById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user entity.Users

	result := database.DB.First(&user, id)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	userResponse := response.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return ctx.JSON(userResponse)
}

func UserUpdateName(ctx *fiber.Ctx) error {
	// Get User by ID

	userRequest := new(request.UserUpdateRequest)

	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	userId := ctx.Params("id")
	var user entity.Users

	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found",
			"data": nil,
		})
	}

	// Update user Name or Phone

	user.Name = userRequest.Name
	if userRequest.Name != "" {
		user.Name = userRequest.Name
	}

	errUpdate := database.DB.Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Could not update user",
		})
	}

	return ctx.JSON(fiber.Map{
		"status": "success",
		"message": "User successfully created",
		"data": user,
	})
}


func UserDelete(ctx *fiber.Ctx) error {
	// get user by id

	id := ctx.Params("id")
	var user entity.Users

	result := database.DB.First(&user, id)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	// delete user

	err := database.DB.Delete(&user, id).Error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Could not delete user",
		})
	}

	return ctx.JSON(fiber.Map{
		"status": "success",
		"message": "User successfully deleted",
	})
}