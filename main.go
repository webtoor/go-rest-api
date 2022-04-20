package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/webtoor/go-fiber/config"
	"github.com/webtoor/go-fiber/controller"
	"github.com/webtoor/go-fiber/exception"
	"github.com/webtoor/go-fiber/repository"
	"github.com/webtoor/go-fiber/service"
)

func main() {
	db := config.InitDb()
	validate := validator.New()
	CoinRepository := repository.NewCoinRepository()
	CoinService := service.NewCoinService(CoinRepository, db, validate)
	CoinController := controller.NewCoinController(CoinService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	CoinController.Route(app)

	// Start App
	err := app.Listen(":3000")
	exception.Panic(err)

}
