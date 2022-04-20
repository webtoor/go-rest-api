package repository

import (
	"github.com/webtoor/go-fiber/config"
	"github.com/webtoor/go-fiber/model/entity"
)

type CoinRepositoryImpl struct {
}

func NewCoinRepository() CoinRepository {
	return &CoinRepositoryImpl{}
}

func (repository *CoinRepositoryImpl) Create(Coin entity.Coin) entity.Coin {

	if result := config.DB.Create(&Coin); result.Error != nil {
		panic(result.Error)
	}

	return Coin
}
