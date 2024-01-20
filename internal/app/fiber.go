package app

import "github.com/gofiber/fiber/v2"

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		AppName:               "go-rest-api",
		DisableStartupMessage: true,
	}
}
