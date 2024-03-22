package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lenglee-vaja/blogbackend/controller"
	"github.com/lenglee-vaja/blogbackend/middleware"
)

func Setup(a *fiber.App) {
	// a.Use(middleware.IsAuthenticated)
	a.Post("/api/register", controller.Register)
	a.Post("/api/login", controller.Login)

	a.Use(middleware.IsAuthenticated)
	a.Post("/api/post", controller.CreatePost)
	a.Get("/api/posts", controller.AllPosts)
	a.Get("/api/post/:id", controller.GetPost)
	a.Put("/api/post/:id", controller.UpdatePost)
	a.Get("/api/unique-post", controller.UniquePost)
	a.Delete("/api/post/:id", controller.DeletePost)
	a.Post("/api/upload-image", controller.UploadImage)
	a.Static("/api/uploads", "./uploads")
}
