package net_worth

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

func (r *repository) CreateNetWorth(ctx context.Context, worth NetWorth) error {
	return r.createNetWorth(ctx, worth)
}

func (r *repository) SetNetWorth(ctx context.Context, worth NetWorth) (*NetWorth, error) {
	return r.setNetWorth(ctx, worth)
}
