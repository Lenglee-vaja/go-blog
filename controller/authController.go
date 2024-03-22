package controller

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/lenglee-vaja/blogbackend/database"
	"github.com/lenglee-vaja/blogbackend/model"
	"github.com/lenglee-vaja/blogbackend/util"
)

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z0-9._%+\-]`)
	return Re.MatchString(email)
}

func Register(c *fiber.Ctx) error {

	var data map[string]interface{}
	var userData model.User
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	//Check if password is less then 6 characters
	if len(data["password"].(string)) < 6 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": "Password must be greater than 6 characters",
		})
	}
	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": "Invalid Email",
		})
	}
	//Check if email already exists
	database.DB.Where("email = ?", strings.TrimSpace(data["email"].(string))).First(&userData)
	if userData.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": "Email already exists",
		})
	}
	user := model.User{
		FirstName: data["first_name"].(string),
		LastName:  data["last_name"].(string),
		Email:     strings.TrimSpace(data["email"].(string)),
		Phone:     data["phone"].(string),
	}
	user.SetPassword(data["password"].(string))
	err := database.DB.Create(&user)
	if err != nil {
		fmt.Println("err: ", err)
	}
	c.Status(200)
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Success",
		"user":    user,
	})
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	var user model.User
	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": "User not found",
		})
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": "Password does not match",
		})
	}
	token, err := util.GenerateJWT(strconv.Itoa(int(user.Id)))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": "Could not login",
		})
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Success",
		"user":    user,
		"token":   token,
	})
}
