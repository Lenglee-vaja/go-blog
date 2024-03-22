package controller

import (
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randletter(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func UploadImage(c *fiber.Ctx) error {

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["image"]
	fileName := ""
	for _, file := range files {
		fileName = randletter(5) + "-" + file.Filename
		if err := c.SaveFile(file, "./uploads/"+fileName); err != nil {
			return nil
		}
	}
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Success",
		"url":     "http://localhost:3000/api/uploads/" + fileName,
	})
}
