package main

import (
	"auth/db"
	"auth/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8080")
}
