package user

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

const userCollectionName = "user"
const authCollectionName = "auth"

func (r *repository) getUserByHandle(ctx context.Context, handle string) (*User, error) {
	userColl := r.db.Collection(userCollectionName)
	var result User
	err := userColl.FindOne(ctx, bson.M{"handle": handle}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repository) getUserAuthByHandle(ctx context.Context, handle string) (*Auth, error) {
	authColl := r.db.Collection(authCollectionName)
	var result Auth
	err := authColl.FindOne(ctx, bson.M{"handle": handle}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repository) createUser(ctx context.Context, user User, auth Auth) error {
	userColl := r.db.Collection(userCollectionName)
	_, err := userColl.InsertOne(ctx, user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			if strings.Contains(err.Error(), "unique_email") {
				return fmt.Errorf("email already exists")
			}

			if strings.Contains(err.Error(), "unique_handle") {
				return fmt.Errorf("user handle already exists")
			}
		}
		return fmt.Errorf("something went wrong")
	}

	authColl := r.db.Collection(authCollectionName)
	_, err = authColl.InsertOne(ctx, auth)
	if err != nil {
		return err
	}
	return nil
}
