package jwt

import (
	"errors"
	"fmt"
	"os"
	"time"

	jwt "github.com/form3tech-oss/jwt-go"
)

const TOKEN_PREFIX = "Bearer"

func GenerateToken(claimsUser map[string]string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	for name, value := range claimsUser {
		claims[name] = value
	}
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	assignedToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("Internal error")
	}
	return fmt.Sprintf("%s %s", TOKEN_PREFIX, assignedToken), nil
}
