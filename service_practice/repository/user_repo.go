package repository

import (
	"go_sample/model"

	"gorm.io/gorm"
)

// レポジトリ層で使用する関数たち
type UserRepository interface {
	FindByID(id uint) (*model.User, error)
	Create(u *model.User) error
}

// 上を実現するために
type userRepo struct {
	db *gorm.DB
}

func (r *userRepo) FindByID(id int) (*model.User, error) {
	var u model.User
	if err := r.db.First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepo) Create(u *model.User) error {
	return r.db.Create(u).Error
}
