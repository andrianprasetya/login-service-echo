package auth

import "login-service/entity"

type Repository interface {
	Login(email string) (entity.User, error)
	Register(dto RegisterDto) (entity.User, error)
}
