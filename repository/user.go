package repository

import (
	"bwastartup/entity"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user entity.User) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) Create(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
