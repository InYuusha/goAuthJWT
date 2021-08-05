package controllers

import (
	"auth/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)

	if err != nil {
		return err
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]),14)
	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	return c.JSON(user)
}
