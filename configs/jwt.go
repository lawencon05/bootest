package configs

import (
	"time"

	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

var (
	UserIdExample    = "20022012"
	SecretKeyExample = "mLmHu8f1IxFo4dWurBG3jEf1Ex0wDZvvwND6eFmcaX"
	SigningMethodExample = "HS512"
)

type JwtClaims struct {
	Name string `json:"name"`
	jwtGo.StandardClaims
}

func SetJwt(e *echo.Echo) (*echo.Group){
	jwtGroup := e.Group("/api")

	jwtGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig {
		SigningMethod: SigningMethodExample,
		SigningKey: [] byte(SecretKeyExample),
	}))

	return jwtGroup
}

func CreateJwtToken(username string) (string, error) {
	claims := JwtClaims{
		username,
		jwtGo.StandardClaims{
			Id:        username,
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		},
	}

	rawToken := jwtGo.NewWithClaims(jwtGo.SigningMethodHS512, claims)
	token, err := rawToken.SignedString([]byte(SecretKeyExample))
	if err != nil {
		return "", err
	}

	return token, nil
}
