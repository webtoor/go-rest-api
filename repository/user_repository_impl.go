package repository

import (
	"github.com/webtoor/go-fiber/config"
	"github.com/webtoor/go-fiber/model/entity"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(user entity.User) entity.User {

	if result := config.DB.Create(&user); result.Error != nil {
		panic(result.Error)
	}

	return user
}

func (repository *UserRepositoryImpl) FindAll() []entity.User {
	users := []entity.User{}

	config.DB.Find(&users)

	return users
}

func (repository *UserRepositoryImpl) FindById(userId int) (entity.User, error) {
	user := entity.User{}

	if result := config.DB.First(&user, userId); result.Error != nil {
		panic(result.Error)
	}

	return user, nil
}

func (repository *UserRepositoryImpl) Update(userId int, user entity.User) entity.User {

	if result := config.DB.Model(&user).Where("id = ?", userId).Updates(&user); result.Error != nil {
		panic(result.Error)
	}

	return user
}

func (repository *UserRepositoryImpl) Delete(userId int) {
	config.DB.Delete(&entity.User{}, userId)
}
