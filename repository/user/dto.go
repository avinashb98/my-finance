package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Handle    string             `bson:"handle,omitempty"`
	Name      string             `bson:"name,omitempty"`
	Email     string             `bson:"email,omitempty"`
	IsActive  bool               `bson:"is_active,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
}
