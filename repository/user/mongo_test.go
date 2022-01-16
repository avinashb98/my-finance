package user_test

import (
	"context"
	"fmt"
	"github.com/avinashb98/myfin/datasources/mongo"
	mocks "github.com/avinashb98/myfin/mocks/datasources/mongo"
	"github.com/avinashb98/myfin/repository/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestRepository_GetUserByHandle(t *testing.T) {
	var db mongo.Database
	var collection mongo.Collection
	var singleResultErr mongo.SingleResult
	var singleResultNotFoundErr mongo.SingleResult
	var singleResultFound mongo.SingleResult

	db = &mocks.Database{}
	collection = &mocks.Collection{}
	singleResultErr = &mocks.SingleResult{}
	singleResultNotFoundErr = &mocks.SingleResult{}
	singleResultFound = &mocks.SingleResult{}

	singleResultErr.(*mocks.SingleResult).
		On("Decode", mock.AnythingOfType("*user.User")).
		Return(fmt.Errorf("mocked-error"))

	singleResultNotFoundErr.(*mocks.SingleResult).
		On("Decode", mock.AnythingOfType("*user.User")).
		Return(fmt.Errorf("mongo: no documents in result"))

	singleResultFound.(*mocks.SingleResult).
		On("Decode", mock.AnythingOfType("*user.User")).
		Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*user.User)
		arg.Handle = "correctUserHandle"
	})

	collection.(*mocks.Collection).
		On("FindOne", context.Background(), bson.M{"handle": "errorUserHandle"}).
		Return(singleResultErr)

	collection.(*mocks.Collection).
		On("FindOne", context.Background(), bson.M{"handle": "notFoundHandle"}).
		Return(singleResultNotFoundErr)

	collection.(*mocks.Collection).
		On("FindOne", context.Background(), bson.M{"handle": "correctUserHandle"}).
		Return(singleResultFound)

	db.(*mocks.Database).
		On("Collection", "user").
		Return(collection)

	repo := user.NewRepository(context.Background(), db)

	userReturned, err := repo.GetUserByHandle(context.Background(), "errorUserHandle")
	assert.Empty(t, userReturned)
	assert.EqualError(t, err, "mocked-error")

	userReturned, err = repo.GetUserByHandle(context.Background(), "notFoundHandle")
	assert.Empty(t, userReturned)
	assert.NotEmpty(t, err)
	assert.Equal(t, err.Error(), "mongo: no documents in result")

	userReturned, err = repo.GetUserByHandle(context.Background(), "correctUserHandle")
	assert.Equal(t, "correctUserHandle", userReturned.Handle)
	assert.NoError(t, err)
}

func TestRepository_CreateUser(t *testing.T) {
	var db mongo.Database
	var collection mongo.Collection

	db = &mocks.Database{}
	collection = &mocks.Collection{}

	collection.(*mocks.Collection).
		On("InsertOne", context.Background(), user.User{Handle: "validHandleName"}).
		Return("", nil)

	collection.(*mocks.Collection).
		On("InsertOne", context.Background(), user.User{Handle: "invalidHandleName"}).
		Return("", fmt.Errorf("invalid handle"))

	db.(*mocks.Database).
		On("Collection", "user").
		Return(collection)

	repo := user.NewRepository(context.Background(), db)
	err := repo.CreateUser(context.Background(), user.User{Handle: "validHandleName"})
	assert.Empty(t, err)
	err = repo.CreateUser(context.Background(), user.User{Handle: "invalidHandleName"})
	assert.NotEmpty(t, err)
	assert.Equal(t, err.Error(), "invalid handle")
}
