package service

import (
	"github.com/webtoor/go-fiber/model/web"
)

type UserService interface {
	Create(request web.UserCreateRequest) (response web.UserCreateResponse)
	FindAll() (response []web.GetUserResponse)
	FindById(userId int) (response web.GetUserResponse)
	Update(userId int, request web.UserUpdateRequest) (response web.UserUpdateResponse)
	Delete(userId int)
}
