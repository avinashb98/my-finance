package net_worth_test

import (
	"context"
	"fmt"
	"github.com/avinashb98/myfin/datasources/mongo"
	mocks "github.com/avinashb98/myfin/mocks/datasources/mongo"
	"github.com/avinashb98/myfin/repository/net_worth"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_CreateNetWorth(t *testing.T) {
	var db mongo.Database
	var netWorthCollection mongo.Collection
	db = &mocks.Database{}
	netWorthCollection = &mocks.Collection{}
	netWorthCollection.(*mocks.Collection).
		On("InsertOne", context.Background(), net_worth.NetWorth{Handle: "validHandleName"}).
		Return("", nil)

	netWorthCollection.(*mocks.Collection).
		On("InsertOne", context.Background(), net_worth.NetWorth{Handle: "inValidHandleName"}).
		Return("", fmt.Errorf("invalid handle"))

	netWorthCollection.(*mocks.Collection).
		On("InsertOne", context.Background(), net_worth.NetWorth{Handle: "duplicateHandleName"}).
		Return("", fmt.Errorf("handle already exists"))

	db.(*mocks.Database).
		On("Collection", "net_worth").
		Return(netWorthCollection)

	repo := net_worth.NewRepository(context.Background(), db)

	err := repo.CreateNetWorth(context.Background(), net_worth.NetWorth{Handle: "validHandleName"})
	assert.Empty(t, err)

	err = repo.CreateNetWorth(context.Background(), net_worth.NetWorth{Handle: "inValidHandleName"})
	assert.NotEmpty(t, err)
	assert.Equal(t, err.Error(), "something went wrong")

	err = repo.CreateNetWorth(context.Background(), net_worth.NetWorth{Handle: "duplicateHandleName"})
	assert.NotEmpty(t, err)
	assert.Equal(t, err.Error(), "something went wrong")

}
