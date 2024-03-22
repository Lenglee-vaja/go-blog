package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lenglee-vaja/blogbackend/controller"
)

func Setup(a *fiber.App) {
	// a.Use(middleware.IsAuthenticated)
	a.Post("/api/register", controller.Register)
	a.Post("/api/login", controller.Login)
}
