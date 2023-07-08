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
		SigningKey:  GetJwtSecretKey(),
	})

func GenerateTokenPair(user entity.User) (*string, *string, interface{}, error) {

	// Create token with claims
	token := jwt.New(jwt.SigningMethodHS256)
	tokenClaims := token.Claims.(jwt.MapClaims)

	tokenClaims["id"] = user.ID
	tokenClaims["email"] = user.Email
	tokenClaims["name"] = user.Name
	tokenClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Generate encoded token and send it as response.
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["email"] = user.Email
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	//Encode Token
	accessToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	//Encode Refresh Token
	rt, err := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		return nil, nil, nil, err
	}

	return &accessToken, &rt, tokenClaims["exp"], nil
}
