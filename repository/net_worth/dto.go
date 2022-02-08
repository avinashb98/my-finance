package net_worth

import (
	"time"
)

type NetWorth struct {
	Handle    string    `bson:"handle,omitempty"`
	NetWorth  int       `bson:"net_worth"`
	UpdatedAt time.Time `bson:"updated_at,omitempty"`
}
