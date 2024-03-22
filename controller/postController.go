package controller

import (
	"fmt"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/lenglee-vaja/blogbackend/database"
	"github.com/lenglee-vaja/blogbackend/model"
	"github.com/lenglee-vaja/blogbackend/util"
)

func CreatePost(c *fiber.Ctx) error {
	var blogPost model.Blog
	if err := c.BodyParser(&blogPost); err != nil {
		fmt.Println("Unable to parse body")
	}
	if err := database.DB.Create(&blogPost).Error; err != nil {
		c.Status(500)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": "Invalid payload",
		})
	}
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Success",
		"id":      blogPost.Id,
	})
}

func AllPosts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := (page - 1) * limit
	var total int64
	var getBlog []model.Blog
	database.DB.Preload("User").Limit(limit).Offset(offset).Order("id desc").Find(&getBlog)
	database.DB.Model(model.Blog{}).Count(&total)
	return c.JSON(fiber.Map{
		"data":   getBlog,
		"status": true,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})
}

func GetPost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var blogPost model.Blog
	database.DB.Preload("User").First(&blogPost, id)
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Success",
		"data":    blogPost,
	})
}

func UpdatePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog := model.Blog{
		Id: uint(id),
	}
	if err := c.BodyParser(&blog); err != nil {
		fmt.Println("Unable to parse body")
	}
	database.DB.Model(&blog).Updates(&blog)
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Success",
		"id":      blog.Id,
	})
}

func UniquePost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := util.ParesJWT(cookie)
	var blog []model.Blog
	database.DB.Where("user_id = ?", id).Preload("User").Find(&blog)
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Success",
		"data":    blog,
	})
}

func DeletePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog := model.Blog{
		Id: uint(id),
	}
	database.DB.Delete(&blog)
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Success",
	})
}
