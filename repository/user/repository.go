package user

import (
	"context"
	"github.com/avinashb98/myfin/datasources/mongo"
)

type repository struct {
	ctx context.Context
	db  mongo.Database
}

func NewRepository(ctx context.Context, db mongo.Database) Repository {
	return &repository{
		ctx: ctx,
		db:  db,
	}
}

func (r *repository) GetUserByHandle(ctx context.Context, handle string) (*User, error) {
	return r.getUserByHandle(ctx, handle)
}

func (r *repository) CreateUser(ctx context.Context, user User) error {
	return r.createUser(ctx, user)
}
