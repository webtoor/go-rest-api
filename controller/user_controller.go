package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/webtoor/go-rest-api/helper"
	"github.com/webtoor/go-rest-api/model/web"
	"github.com/webtoor/go-rest-api/service"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(UserService service.UserService) UserController {
	return UserController{
		UserService: UserService,
	}
}

func (controller *UserController) Route(app *fiber.App) {
	app.Post("/users", controller.Create)
	app.Get("/users", controller.FindAll)
	app.Get("/users/:id", controller.FindById)
	app.Put("/users/:id", controller.Update)
	app.Delete("/users/:id", controller.Delete)
}

func (controller *UserController) Create(c *fiber.Ctx) error {
	var request web.UserCreateRequest
	err := c.BodyParser(&request)
	helper.PanicIfError(err)

	response := controller.UserService.Create(request)
	return c.JSON(web.JsonResponse{
		Code:   201,
		Status: "CREATED",
		Data:   response,
	})
}

func (controller *UserController) FindAll(c *fiber.Ctx) error {
	response := controller.UserService.FindAll()
	return c.JSON(web.JsonResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *UserController) FindById(c *fiber.Ctx) error {
	userId := c.Params("id")
	id, _ := strconv.Atoi(userId)
	response := controller.UserService.FindById(id)
	return c.JSON(web.JsonResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *UserController) Update(c *fiber.Ctx) error {
	var request web.UserUpdateRequest
	err := c.BodyParser(&request)
	helper.PanicIfError(err)

	userId := c.Params("id")
	id, _ := strconv.Atoi(userId)
	response := controller.UserService.Update(id, request)
	return c.JSON(web.JsonResponse{
		Code:   200,
		Status: "UPDATED",
		Data:   response,
	})
}

func (controller *UserController) Delete(c *fiber.Ctx) error {
	userId := c.Params("id")
	id, _ := strconv.Atoi(userId)
	controller.UserService.Delete(id)
	return c.JSON(web.JsonResponse{
		Code:   200,
		Status: "OK",
	})
}
