package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/webtoor/go-fiber/exception"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
