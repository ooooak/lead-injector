package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ooooak/shop/auth/db"
)

func Login(c *fiber.Ctx) error {
	user := db.Auth{Email: "test@gmail.com"}
	db.DB.Create(user)
	return c.JSON(user)
}
