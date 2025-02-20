package handler

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	"restapi-fiber-gorm/database"
	"restapi-fiber-gorm/model/entity"
	"restapi-fiber-gorm/model/request"
	// "restapi-fiber-gorm/model/response"
	// "restapi-fiber-gorm/utils"

	"github.com/go-playground/validator/v10"
)

func CreateBook(ctx *fiber.Ctx) error {
	book := new(request.BookCreateRequest)

	if err := ctx.BodyParser(book); err != nil {
		return err
	}

	// Request Validation
	validate := validator.New()
	errValidate := validate.Struct(book)
	if errValidate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": "Validation error",
			"error": errValidate.Error(),
		})
	}

	//	Validation Required Image
	var filenameString string

	filename := ctx.Locals("filename")
	log.Println("filename = ", filename)
	if filename == nil {
		return ctx.Status(422).JSON(fiber.Map{
			"message": "image cover is required.",
		})
	} else {
		filenameString = fmt.Sprintf("%v", filename)
	}

	newBook := entity.Books{
		Title:  book.Title,
		Author: book.Author,
		Cover:  filenameString,
	}

	errCreateBook := database.DB.Create(&newBook).Error
	if errCreateBook!= nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Could not create book", 
			"data": nil,
		})
	}

	return ctx.JSON(fiber.Map{
		"status": "success",
		"message": "User successfully created",
		"data": newBook,
	})

}
