package user

import (
	"time"
)

type User struct {
	Handle    string    `bson:"handle,omitempty"`
	Name      string    `bson:"name,omitempty"`
	Email     string    `bson:"email,omitempty"`
	IsActive  bool      `bson:"is_active,omitempty"`
	CreatedAt time.Time `bson:"created_at,omitempty"`
}

type Auth struct {
	Handle       string    `bson:"handle"`
	PasswordHash string    `bson:"password_hash"`
	LastLogin    time.Time `bson:"last_login"`
}
