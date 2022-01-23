package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Payload struct {
	Handle string `json:"handle"`
	User   bool   `json:"user"`
	jwt.StandardClaims
}

func (s *service) GenerateToken(handle string, isUser bool) (string, error) {
	claims := &Payload{
		handle,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(s.config.JWTUserTokenExpiryInHours)).Unix(),
			Issuer:    s.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (s *service) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(encodedToken, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})
}
