package controllers

import (
	"auth/db"
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
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}
	db.DB.Create(&user)

	return c.JSON(user)
}
func Login(c *fiber.Ctx)error{

	var data map[string]string
	err := c.BodyParser(&data)

	if err != nil {
		return err
	}
	var user models.User
	db.DB.Where("email=?",data["email"]).First(&user)

	if user.Id==0{
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"Message":"User not found",
		})
	}
	if err:=bcrypt.CompareHashAndPassword(user.Password,[]byte(data["password"]));err!=nil{
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"Message":"Incorrect Password",
		})
	}
	return c.JSON(user)


}
