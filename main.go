package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	authHandler "login-service/auth/delivery/http"
	authService "login-service/auth/usecase"
	"login-service/infrastructure/database"
	jwtMiddleware "login-service/middleware"
	userHandler "login-service/user/delivery/http"
	userService "login-service/user/usecase"
	"os"
	"path/filepath"
)

func init() {
	fileExecutable, _ := os.Executable()
	basePath, _ := filepath.Split(fileExecutable)
	if os.Getenv("APP_ENV") != "production" {
		basePath = ""
	}
	_ = godotenv.Load(basePath + ".env")
}

func main() {
	e := echo.New()
	db := database.New()
	database.AutoMigrate(db)
	e.Logger.SetLevel(log.DEBUG)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	api := e.Group("/api")
	v1 := api.Group("/v1")
	//AuthController
	authController := authHandler.NewAuthController(authService.NewAuthService(db))
	auth := v1.Group("/auth")
	auth.POST("/token", authController.Login)
	auth.POST("/register", authController.Register)

	//UserController
	userController := userHandler.NewUserController(userService.NewUserService(db))
	user := v1.Group("/user")
	user.GET("/:id", userController.FindById, jwtMiddleware.IsLoggedIn)

	e.Logger.Fatal(e.Start(os.Getenv("APP_PORT")))

}
