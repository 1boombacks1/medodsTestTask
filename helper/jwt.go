package helper

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/google/uuid"
)

type AuthClaims struct {
	jwt.StandardClaims
	Guid string `json:"guid"`
}

func MakeJWT(claims AuthClaims, signKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	AccessToken, err := token.SignedString([]byte(signKey))
	if err != nil {
		return "", fmt.Errorf("token.SignedString: %w", err)
	}

	return AccessToken, nil
}

func NewRefreshToken() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
