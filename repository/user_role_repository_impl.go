package repository

import (
	"github.com/webtoor/go-rest-api/helper"
	"github.com/webtoor/go-rest-api/model/entity"
	"gorm.io/gorm"
)

type UserRoleRepositoryImpl struct {
}

func NewUserRoleRepository() UserRoleRepository {
	return &UserRoleRepositoryImpl{}
}

func (repository *UserRoleRepositoryImpl) Create(tx *gorm.DB, userRole entity.UserRole) entity.UserRole {

	err := tx.Create(&userRole).Error
	helper.PanicIfError(err)
	return userRole
}
