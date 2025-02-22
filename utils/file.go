package utils

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

const DefaultPathAssetImage = "./public/covers/"

func HandleSingleFile(ctx *fiber.Ctx) error {
	//	HANDLE FILE
	file, errFile := ctx.FormFile("cover")
	if errFile != nil {
		log.Println("Error File = ", errFile)
	}

	var filename *string
	if file != nil {
		errCheckContentType := checkContentType(file, "image/jpg", "image/png", "image/gif")
		if errCheckContentType != nil {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"message": errCheckContentType.Error(),
			})
		}

		filename = &file.Filename
		extenstionFile := filepath.Ext(*filename)
		newFilename := fmt.Sprintf("gambar-satu%s", extenstionFile)

		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/covers/%s", newFilename))
		if errSaveFile != nil {
			log.Println("Fail to store file into public/covers directory.")
		}
	} else {
		log.Println("Nothing file to uploading.")
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

func HandleRemoveFile(filename string, pathFile ...string) error {
	var filePath string

	// Tentukan path file
	if len(pathFile) > 0 {
		filePath = filepath.Join(pathFile[0], filename)
	} else {
		filePath = filepath.Join(DefaultPathAssetImage, filename)
	}

	// Cek apakah file ada sebelum dihapus
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		log.Println("File not found:", filePath)
		return fmt.Errorf("file not found: %s", filePath)
	}

	// Hapus file
	err := os.Remove(filePath)
	if err != nil {
		log.Println("Failed to remove file:", err)
		return err
	}

	log.Println("File successfully deleted:", filePath)
	return nil
}


func checkContentType(file *multipart.FileHeader, contentTypes ...string) error {
	if len(contentTypes) > 0 {
		for _, contentType := range contentTypes {
			contentTypeFile := file.Header.Get("Content-Type")
			if contentTypeFile == contentType {
				return nil
			}
		}
		return errors.New("not allowed file type")
	} else {
		return errors.New("not found content type to be checking")
	}
}