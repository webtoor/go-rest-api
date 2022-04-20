package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/webtoor/go-fiber/exception"
	"github.com/webtoor/go-fiber/model/entity"
	"github.com/webtoor/go-fiber/model/web"
	"github.com/webtoor/go-fiber/repository"
	"gorm.io/gorm"
)

var validate = validator.New()

type CoinServiceImpl struct {
	CoinRepository repository.CoinRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewCoinService(CoinRepository repository.CoinRepository, DB *gorm.DB, validate *validator.Validate) CoinService {
	return &CoinServiceImpl{
		CoinRepository: CoinRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *CoinServiceImpl) Create(request web.CoinCreateRequest) web.CoinCreateResponse {

	err := service.Validate.Struct(request)
	exception.Panic(err)

	coin := entity.Coin{
		Coinname: request.Coinname,
	}

	coin = service.CoinRepository.Create(coin)

	response := web.CoinCreateResponse{
		ID:       coin.ID,
		Coinname: coin.Coinname,
	}

	return response
}
