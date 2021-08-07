package controllers

import (
	"auth/db"
	"auth/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const secretKey = "Key"

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
func Login(c *fiber.Ctx) error {

	var data map[string]string
	err := c.BodyParser(&data)

	if err != nil {
		return err
	}
	var user models.User
	db.DB.Where("email=?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"Message": "User not found",
		})
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"Message": "Incorrect Password",
		})
	}
	//token
	clams := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.StandardClaims{
		Issuer:strconv.Itoa(int(user.Id)),
		ExpiresAt:time.Now().Add(time.Hour * 24).Unix(),
	})
	token,err:=clams.SignedString([]byte(secretKey))
	if err!=nil{
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message":"Could not login",
		})
	}
	cookie:=fiber.Cookie{
		Name:"jwt",
		Value:token,
		Expires: time.Now().Add(time.Hour * 24),
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message":"Success",
	})

}
