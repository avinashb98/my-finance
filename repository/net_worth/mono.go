package net_worth

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
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

func (r *repository) getNetWorth(ctx context.Context, handle string) (*NetWorth, error) {
	netWorthColl := r.db.Collection(NetWorthCollectionName)
	var result NetWorth
	err := netWorthColl.FindOne(ctx, bson.M{"handle": handle}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user net worth not found")
		}
		return nil, fmt.Errorf("something went wrong")
	}
	return &result, nil
}

func (r *repository) updateNetWorth(ctx context.Context, netWorth NetWorth) (*NetWorth, error) {
	netWorthColl := r.db.Collection(NetWorthCollectionName)
	var result NetWorth
	err := netWorthColl.FindOneAndUpdate(
		ctx,
		bson.M{"handle": netWorth.Handle},
		bson.M{"$set": bson.M{
			"net_worth":  netWorth.NetWorth,
			"updated_at": netWorth.UpdatedAt,
		}},
	).Decode(&result)

	if err != nil {
		return nil, fmt.Errorf("error while setting net worth")
	}
	result.NetWorth = netWorth.NetWorth
	return &result, nil
}

func (r *repository) setNetWorth(ctx context.Context, netWorth NetWorth) (*NetWorth, error) {
	exists, err := r.exists(ctx, netWorth.Handle)
	if err != nil {
		return nil, err
	}

	if !exists {
		err := r.createNetWorth(ctx, netWorth)
		if err != nil {
			return nil, err
		}
		return r.getNetWorth(ctx, netWorth.Handle)
	}

	return r.updateNetWorth(ctx, netWorth)
}

func (r *repository) exists(ctx context.Context, handle string) (bool, error) {
	netWorthColl := r.db.Collection(NetWorthCollectionName)
	var result NetWorth
	err := netWorthColl.FindOne(ctx, bson.M{"handle": handle}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, fmt.Errorf("something went wrong")
	}
	return true, nil
}
