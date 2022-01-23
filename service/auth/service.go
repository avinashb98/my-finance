package auth

import (
	"context"
	"github.com/avinashb98/myfin/config"
	"github.com/avinashb98/myfin/repository/user"
	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(string, bool) (string, error)
	ValidateToken(string) (*jwt.Token, error)
	IsAuthenticated(context.Context, string, string) (bool, error)
}

type service struct {
	secretKey string
	issuer    string
	userRepo  user.Repository
	config    config.JWT
}

func NewService(conf config.JWT, userRepo user.Repository) Service {
	return &service{
		secretKey: conf.Secret,
		issuer:    conf.Issuer,
		config:    conf,
		userRepo:  userRepo,
	}
}
