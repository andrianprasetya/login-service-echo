package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4/middleware"
	"login-service/entity"
	"os"
	"path/filepath"
	"time"
)

func GetJwtSecretKey() []byte {
	fileExecutable, _ := os.Executable()
	basepath, _ := filepath.Split(fileExecutable)
	if os.Getenv("APP_ENV") != "production" {
		basepath = ""
	}
	_ = godotenv.Load(basepath + ".env")
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	return []byte(jwtSecretKey)
}

var IsLoggedIn = middleware.JWTWithConfig(
	middleware.JWTConfig{
		SigningKey: GetJwtSecretKey(),
	})

func GenerateTokenPair(user entity.User) (*string, error) {
	// Create token with claims
	// Set custom claims
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["name"] = user.Name
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	//Encode Token
	accessToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		return nil, err
	}

	return &accessToken, nil
}
