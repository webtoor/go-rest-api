package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/webtoor/go-fiber/exception"
	"github.com/webtoor/go-fiber/model/web"
	"github.com/webtoor/go-fiber/service"
)

type CoinController struct {
	CoinService service.CoinService
}

func NewCoinController(CoinService service.CoinService) CoinController {
	return CoinController{
		CoinService: CoinService,
	}
}

func (controller *CoinController) Route(app *fiber.App) {
	app.Post("/coins", controller.Create)
}

func (controller *CoinController) Create(c *fiber.Ctx) error {
	var request web.CoinCreateRequest
	err := c.BodyParser(&request)
	exception.Panic(err)

	response := controller.CoinService.Create(request)
	return c.JSON(web.JsonResponse{
		Code:   201,
		Status: "OK",
		Data:   response,
	})
}
