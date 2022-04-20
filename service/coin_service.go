package service

import (
	"github.com/webtoor/go-fiber/model/web"
)

type CoinService interface {
	Create(request web.CoinCreateRequest) (response web.CoinCreateResponse)
}
