package usecase

import (
	"gorm.io/gorm"
	"login-service/entity"
	"login-service/user"
)

type UserService struct {
	*gorm.DB
}

func NewUserService(db *gorm.DB) user.Repository {
	return UserService{db}
}



func (u UserService) FindById(id string) (*entity.User, error) {
	var model entity.User
	model.ID = id
	err := u.DB.First(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, err
}


