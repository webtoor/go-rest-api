package repository

import (
	"github.com/webtoor/go-fiber/model/entity"
)

type UserRepository interface {
	Create(user entity.User) entity.User
	FindAll() []entity.User
	FindById(userId int) (entity.User, error)
	Update(userId int, user entity.User) entity.User
	Delete(userId int)
}
