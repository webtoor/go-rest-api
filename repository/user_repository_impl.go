package repository

import (
	"github.com/webtoor/go-fiber/helper"
	"github.com/webtoor/go-fiber/model/entity"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(tx *gorm.DB, user entity.User) entity.User {

	err := tx.Create(&user).Error
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) FindAll(tx *gorm.DB) []entity.User {
	users := []entity.User{}

	err := tx.Order("id desc").Find(&users).Error
	helper.PanicIfError(err)

	return users
}

func (repository *UserRepositoryImpl) FindById(tx *gorm.DB, userId int) (entity.User, error) {
	user := entity.User{}

	err := tx.First(&user, userId).Error
	helper.PanicIfError(err)

	return user, nil
}

func (repository *UserRepositoryImpl) Update(tx *gorm.DB, userId int, user entity.User) entity.User {

	err := tx.Model(&user).Where("id = ?", userId).Updates(&user).Error
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) Delete(tx *gorm.DB, userId int) {
	err := tx.Delete(&entity.User{}, userId).Error
	helper.PanicIfError(err)
}
