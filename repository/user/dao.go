package user

import "context"

type Repository interface {
	GetUserByHandle(context.Context, string) (*User, error)
	CreateUser(context.Context, User) error
}
