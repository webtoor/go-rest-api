package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/webtoor/go-rest-api/config"
	"github.com/webtoor/go-rest-api/controller"
	"github.com/webtoor/go-rest-api/helper"
	"github.com/webtoor/go-rest-api/repository"
	"github.com/webtoor/go-rest-api/service"
)

func init() {
	err := godotenv.Load()
	helper.PanicIfError(err)
}

func main() {
	db := config.InitDb()
	validate := validator.New()

	userRepository := repository.NewUserRepository()
	userRoleRepository := repository.NewUserRoleRepository()

	userService := service.NewUserService(&userRepository, &userRoleRepository, db, validate)
	userController := controller.NewUserController(userService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	userController.Route(app)

	// Start App
	err := app.Listen(":3000")
	helper.PanicIfError(err)

}
