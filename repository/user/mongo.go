package user

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

const collectionName = "user"

func (r *repository) getUserByHandle(ctx context.Context, handle string) (*User, error) {
	userColl := r.db.Collection(collectionName)
	var result User
	err := userColl.FindOne(ctx, bson.M{"handle": handle}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repository) createUser(ctx context.Context, user User) error {
	userColl := r.db.Collection(collectionName)
	_, err := userColl.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
