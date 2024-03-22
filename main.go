package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/lenglee-vaja/blogbackend/database"
	"github.com/lenglee-vaja/blogbackend/routes"
)

func main() {
	database.Connect()
	port := os.Getenv("PORT")
	app := fiber.New()
	routes.Setup(app)
	app.Listen(port)
}
