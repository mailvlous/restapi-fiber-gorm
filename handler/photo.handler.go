package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"restapi-fiber-gorm/database"
	"restapi-fiber-gorm/model/entity"
	"restapi-fiber-gorm/model/request"
	"log"
)

func PhotoHandlerCreate(ctx *fiber.Ctx) error {
	photo := new(request.PhotoCreateRequest)
	if err := ctx.BodyParser(photo); err != nil {
		return err
	}

	// VALIDASI REQUEST
	validate := validator.New()
	errValidate := validate.Struct(photo)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	filenames := ctx.Locals("filenames")
	if filenames == nil {
		return ctx.Status(422).JSON(fiber.Map{
			"message": "image cover is required.",
		})
	} else {
		filenamesData := filenames.([]string)
		for _, filename := range filenamesData {
			newPhoto := entity.Photos{
				Image:      filename,
				CategoryID: photo.CategoryId,
			}

			errCreatePhoto := database.DB.Create(&newPhoto).Error
			if errCreatePhoto != nil {
				log.Println("Some data not saved properly.")
			}
		}
	}

	return ctx.JSON(fiber.Map{
		"messaage": "success",
	})
}