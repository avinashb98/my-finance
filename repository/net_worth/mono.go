package net_worth

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

const NetWorthCollectionName = "net_worth"

func (r *repository) createNetWorth(ctx context.Context, netWorth NetWorth) error {
	netWorthColl := r.db.Collection(NetWorthCollectionName)
	_, err := netWorthColl.InsertOne(ctx, netWorth)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			if strings.Contains(err.Error(), "unique_handle") {
				return fmt.Errorf("user networth already exists")
			}
		}
		return fmt.Errorf("something went wrong")
	}
	return nil
}
