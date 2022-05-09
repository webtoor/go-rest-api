package repository

import (
	"github.com/webtoor/go-rest-api/model/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(tx *gorm.DB, user entity.User) entity.User
	FindAll(tx *gorm.DB) []entity.User
	FindById(tx *gorm.DB, userId int) (entity.User, error)
	Update(tx *gorm.DB, userId int, user entity.User) entity.User
	Delete(tx *gorm.DB, userId int)
}
