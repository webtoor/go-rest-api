package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/webtoor/go-rest-api/internal/app"
)

type router struct {
	cfg *app.Config
	rtr fiber.Router
}

func NewRouter(cfg *app.Config, rtr fiber.Router) Router {
	return &router{
		cfg: cfg,
		rtr: rtr,
	}
}

func (r *router) Route() {
	// define controller

}
