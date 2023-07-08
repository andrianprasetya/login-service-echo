package user

import "login-service/entity"

type Repository interface {
	FindById(id string) (*entity.User, error)
}