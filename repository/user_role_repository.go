package repository

import (
	"github.com/webtoor/go-rest-api/model/entity"
	"gorm.io/gorm"
)

type UserRoleRepository interface {
	Create(tx *gorm.DB, userRole entity.UserRole) entity.UserRole
}
