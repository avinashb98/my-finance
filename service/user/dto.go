package user

import "time"

type User struct {
	Handle    string    `json:"handle"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

type Auth struct {
	Handle       string
	PasswordHash string
}
