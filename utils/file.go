package utils

import (
	"fmt"
	"log"
	"path/filepath"
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

func HandleMultipleFile(ctx *fiber.Ctx) error {
	form, errForm := ctx.MultipartForm()
	if errForm != nil {
		log.Println("Error Read Multipart Form Request, Error = ", errForm)
	}

	files := form.File["photos"]

	var filenames []string

	for i, file := range files {
		var filename string
		if file != nil {
			extenstionFile := filepath.Ext(file.Filename)
			filename = fmt.Sprintf("%d-%s%s", i, "gambar", extenstionFile)

			errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/covers/%s", filename))
			if errSaveFile != nil {
				log.Println("Fail to store file into public/covers directory.")
			}
		} else {
			log.Println("Nothing file to uploading.")
		}

		if filename != "" {
			filenames = append(filenames, filename)
		}
	}

	ctx.Locals("filenames", filenames)

	return ctx.Next()
}
