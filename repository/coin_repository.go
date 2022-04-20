package repository

import (
	"github.com/webtoor/go-fiber/model/entity"
)

type CoinRepository interface {
	Create(Coin entity.Coin) entity.Coin
}
