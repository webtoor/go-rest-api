package repository

import (
	"github.com/webtoor/go-fiber/model/entity"
	"gorm.io/gorm"
)

type UserRoleRepository interface {
	Create(tx *gorm.DB, userRole entity.UserRole) entity.UserRole
}
