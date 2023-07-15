package http

import (
	"github.com/labstack/echo/v4"
	"login-service/auth"
	"login-service/infrastructure/response"
	"login-service/middleware"
	"login-service/utils"
)

type authController struct {
	authRepository auth.Repository
	authMapper     *auth.Mapper
}

func NewAuthController(s auth.Repository) *authController {
	return &authController{authRepository: s,
		authMapper: auth.NewAuthMapper(),
	}
}

func (c *authController) Login(ctx echo.Context) error {
	var dto auth.LoginDto

	if err := ctx.Bind(&dto); err != nil {
		return response.InternalServerError(ctx, utils.InternalServerError, nil, err.Error())
	}
	result, err := c.authRepository.Login(dto.Email)
	if err != nil {
		return response.InternalServerError(ctx, utils.InternalServerError, nil, err.Error())
	}
	if !utils.CheckPasswordHash(dto.Password, result.Password) {
		return response.BadRequest(ctx, utils.BadRequest, nil, "Wrong username or password")
	}
	tokens, err := middleware.GenerateTokenPair(result)
	if err != nil {
		return response.InternalServerError(ctx, utils.InternalServerError, nil, err.Error())
	}
	return response.SingleData(ctx, utils.OK, echo.Map{"access_token": tokens},
		nil)

}

func (c *authController) Register(ctx echo.Context) error {
	var dto auth.RegisterDto
	if err := ctx.Bind(&dto); err != nil {
		return response.InternalServerError(ctx, utils.InternalServerError, nil, err.Error())
	}

	result, err := c.authRepository.Register(dto)
	if err != nil {
		return response.InternalServerError(ctx, utils.InternalServerError, nil, err.Error())
	}
	return response.SingleData(ctx, utils.OK, c.authMapper.Map(result), nil)
}
