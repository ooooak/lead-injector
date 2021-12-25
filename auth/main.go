package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ooooak/shop/auth/db"
	"github.com/ooooak/shop/auth/services"
)

func main() {
	err := db.Connect()
	if err != nil {
		panic(err)
	}

	err = db.Migrate(db.DB)
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	app.Get("/auth/users/create", services.Create)
	app.Get("/auth/users/current", services.Current)
	app.Get("/auth/users/login", services.Login)
	app.Get("/auth/users/login-out", services.Current)

	app.Listen(":3000")
}
