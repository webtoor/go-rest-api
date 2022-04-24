package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/webtoor/go-fiber/exception"
	"github.com/webtoor/go-fiber/model/entity"
	"github.com/webtoor/go-fiber/model/web"
	"github.com/webtoor/go-fiber/repository"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *gorm.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Create(request web.UserCreateRequest) web.UserCreateResponse {

	err := service.Validate.Struct(request)
	exception.Panic(err)

	user := entity.User{
		Email:    request.Email,
		Password: request.Password,
	}

	user = service.UserRepository.Create(user)

	response := web.UserCreateResponse{
		Id:    user.Id,
		Email: user.Email,
	}

	return response
}

func (service *UserServiceImpl) FindAll() (responses []web.GetUserResponse) {

	users := service.UserRepository.FindAll()

	for _, User := range users {
		responses = append(responses, web.GetUserResponse{
			Id:    User.Id,
			Email: User.Email,
		})
	}

	return responses
}

func (service *UserServiceImpl) FindById(userId int) web.GetUserResponse {

	user, err := service.UserRepository.FindById(userId)
	exception.Panic(err)

	response := web.GetUserResponse{
		Id:    user.Id,
		Email: user.Email,
	}

	return response
}

func (service *UserServiceImpl) Update(userId int, request web.UserUpdateRequest) web.UserUpdateResponse {

	err := service.Validate.Struct(request)
	exception.Panic(err)

	user := entity.User{
		Email: request.Email,
	}

	user = service.UserRepository.Update(userId, user)

	response := web.UserUpdateResponse{
		Email: user.Email,
	}

	return response
}

func (service *UserServiceImpl) Delete(userId int) {

	_, err := service.UserRepository.FindById(userId)

	exception.Panic(err)

	service.UserRepository.Delete(userId)
}
