package http

import (
	"github.com/labstack/echo/v4"
	"login-service/infrastructure/response"
	"login-service/user"
	"login-service/utils"
)

type userController struct {
	userRepository user.Repository
	userMapper     *user.Mapper
}

func NewUserController(s user.Repository) *userController {
	return &userController{userRepository: s,
		userMapper: user.NewUserMapper(),
	}
}


func (c *userController) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	result, err := c.userRepository.FindById(id)
	if err != nil {
		return response.InternalServerError(ctx, utils.InternalServerError, nil, err.Error())
	}
	return response.SingleData(ctx, utils.OK, c.userMapper.Map(*result), nil)
}
