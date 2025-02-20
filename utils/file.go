package utils

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func HandleSingleFile(ctx *fiber.Ctx) error {
	// HANDLE FILE
	file, errFile := ctx.FormFile("cover")
	if errFile != nil {
		log.Println("Error file", errFile)
	}

	var filename *string
	if file != nil {
		filename = &file.Filename

		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/covers/%s", *filename))
		if errSaveFile != nil {
			log.Println("Failed")
		}
	} else {
		log.Println("No file")
	}

	if filename != nil {
		ctx.Locals("filename", *filename)
	} else {
		ctx.Locals("filename", nil)
	}

	return ctx.Next()
}
