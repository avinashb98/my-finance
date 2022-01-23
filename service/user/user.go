package user

import (
	"context"
	"github.com/avinashb98/myfin/repository/user"
	"github.com/avinashb98/myfin/utils"
	"time"
)

type Service interface {
	GetUserByHandle(context.Context, string) (*User, error)
	CreateUser(context.Context, User, string) error
	GetUserAuthByHandle(context.Context, string) (*Auth, error)
}

type service struct {
	userRepo user.Repository
}

func NewService(userRepo user.Repository) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (s *service) GetUserByHandle(ctx context.Context, handle string) (*User, error) {
	_user, err := s.userRepo.GetUserByHandle(ctx, handle)
	if err != nil {
		return nil, err
	}

	userDetails := User{
		Handle:    _user.Handle,
		Name:      _user.Name,
		Email:     _user.Email,
		IsActive:  _user.IsActive,
		CreatedAt: _user.CreatedAt,
	}
	return &userDetails, nil
}

func (s *service) GetUserAuthByHandle(ctx context.Context, handle string) (*Auth, error) {
	userAuth, err := s.userRepo.GetUserAuthByHandle(ctx, handle)
	if err != nil {
		return nil, err
	}

	userAuthDetails := Auth{
		Handle:       userAuth.Handle,
		PasswordHash: userAuth.PasswordHash,
	}
	return &userAuthDetails, nil
}

func (s *service) CreateUser(ctx context.Context, userInput User, password string) error {
	passwordHash, err := utils.HashPassword(password, 14)
	if err != nil {
		return err
	}

	_user := user.User{
		Handle:    userInput.Handle,
		Name:      userInput.Name,
		Email:     userInput.Email,
		IsActive:  true,
		CreatedAt: time.Now(),
	}

	_auth := user.Auth{
		Handle:       userInput.Handle,
		PasswordHash: passwordHash,
		LastLogin:    time.Now(),
	}

	return s.userRepo.CreateUser(ctx, _user, _auth)
}
