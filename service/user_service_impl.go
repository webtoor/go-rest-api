package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/webtoor/go-fiber/helper"
	"github.com/webtoor/go-fiber/model/entity"
	"github.com/webtoor/go-fiber/model/web"
	"github.com/webtoor/go-fiber/repository"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository     repository.UserRepository
	UserRoleRepository repository.UserRoleRepository
	DB                 *gorm.DB
	Validate           *validator.Validate
}

func NewUserService(userRepository *repository.UserRepository, userRoleRepository *repository.UserRoleRepository, DB *gorm.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository:     *userRepository,
		UserRoleRepository: *userRoleRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *UserServiceImpl) Create(request web.UserCreateRequest) web.UserCreateResponse {

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// Start DB Transaction
	tx := service.DB.Begin()

	defer helper.CommitOrRollback(tx)

	user := entity.User{
		Email:    request.Email,
		Password: request.Password,
	}

	user = service.UserRepository.Create(tx, user)

	user_role := entity.UserRole{
		UserId: user.Id,
		RoleId: 1,
	}

	user_role = service.UserRoleRepository.Create(tx, user_role)

	response := web.UserCreateResponse{
		Id:    user.Id,
		Email: user.Email,
	}

	return response
}

func (service *UserServiceImpl) FindAll() (responses []web.GetUserResponse) {

	tx := service.DB.Begin()

	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.FindAll(tx)

	for _, User := range users {
		responses = append(responses, web.GetUserResponse{
			Id:    User.Id,
			Email: User.Email,
		})
	}

	return responses
}

func (service *UserServiceImpl) FindById(userId int) web.GetUserResponse {

	tx := service.DB.Begin()

	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(tx, userId)
	helper.PanicIfError(err)

	response := web.GetUserResponse{
		Id:    user.Id,
		Email: user.Email,
	}

	return response
}

func (service *UserServiceImpl) Update(userId int, request web.UserUpdateRequest) web.UserUpdateResponse {

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()

	defer helper.CommitOrRollback(tx)

	service.UserRepository.FindById(tx, userId)

	user := entity.User{
		Email: request.Email,
	}

	user = service.UserRepository.Update(tx, userId, user)

	response := web.UserUpdateResponse{
		Email: user.Email,
	}

	return response
}

func (service *UserServiceImpl) Delete(userId int) {

	tx := service.DB.Begin()

	defer helper.CommitOrRollback(tx)

	_, err := service.UserRepository.FindById(tx, userId)

	helper.PanicIfError(err)

	service.UserRepository.Delete(tx, userId)

}
