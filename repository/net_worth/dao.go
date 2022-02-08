package net_worth

import "context"

type Repository interface {
	CreateNetWorth(context.Context, NetWorth) error
}
