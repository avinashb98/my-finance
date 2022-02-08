package net_worth

import (
	"context"
	"fmt"
	"github.com/avinashb98/myfin/datasources/mongo"
	mocks "github.com/avinashb98/myfin/mocks/datasources/mongo"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	mdb "go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"testing"
)

func Test_repository_createNetWorth(t *testing.T) {
	bctx := context.Background()
	var db mongo.Database
	var netWorthCollection mongo.Collection

	db = &mocks.Database{}
	netWorthCollection = &mocks.Collection{}

	netWorthCollection.(*mocks.Collection).
		On("InsertOne", context.Background(), NetWorth{Handle: "validHandleName"}).
		Return("", nil)

	netWorthCollection.(*mocks.Collection).
		On("InsertOne", context.Background(), NetWorth{Handle: "errorHandleName"}).
		Return("", fmt.Errorf("something went wrong"))

	netWorthCollection.(*mocks.Collection).
		On("InsertOne", context.Background(), NetWorth{Handle: "duplicateHandleName"}).
		Return("", fmt.Errorf("user networth already exists"))

	db.(*mocks.Database).
		On("Collection", "net_worth").
		Return(netWorthCollection)

	type fields struct {
		ctx context.Context
		db  mongo.Database
	}
	type args struct {
		ctx      context.Context
		netWorth NetWorth
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"erroneous user handle",
			fields{
				ctx: bctx,
				db:  db,
			},
			args{bctx, NetWorth{Handle: "errorHandleName"}},
			true,
		},
		{
			"duplicate handle",
			fields{
				ctx: bctx,
				db:  db,
			},
			args{bctx, NetWorth{Handle: "duplicateHandleName"}},
			true,
		},

		{
			"valid handle",
			fields{
				ctx: bctx,
				db:  db,
			},
			args{bctx, NetWorth{Handle: "validHandleName"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				ctx: tt.fields.ctx,
				db:  tt.fields.db,
			}
			if err := r.createNetWorth(tt.args.ctx, tt.args.netWorth); (err != nil) != tt.wantErr {
				t.Errorf("createNetWorth() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_repository_exists(t *testing.T) {
	bctx := context.Background()
	var db mongo.Database
	var netWorthCollection mongo.Collection
	var singleResultErr mongo.SingleResult
	var singleResultNotFoundErr mongo.SingleResult
	var singleResultFound mongo.SingleResult

	db = &mocks.Database{}
	netWorthCollection = &mocks.Collection{}
	singleResultErr = &mocks.SingleResult{}
	singleResultNotFoundErr = &mocks.SingleResult{}
	singleResultFound = &mocks.SingleResult{}

	singleResultErr.(*mocks.SingleResult).
		On("Decode", mock.AnythingOfType("*net_worth.NetWorth")).
		Return(fmt.Errorf("mocked-error"))

	singleResultNotFoundErr.(*mocks.SingleResult).
		On("Decode", mock.AnythingOfType("*net_worth.NetWorth")).
		Return(mdb.ErrNoDocuments)

	singleResultFound.(*mocks.SingleResult).
		On("Decode", mock.AnythingOfType("*net_worth.NetWorth")).
		Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*NetWorth)
		arg.Handle = "correctUserHandle"
	})

	netWorthCollection.(*mocks.Collection).
		On("FindOne", context.Background(), bson.M{"handle": "errorUserHandle"}).
		Return(singleResultErr)

	netWorthCollection.(*mocks.Collection).
		On("FindOne", context.Background(), bson.M{"handle": "notFoundHandle"}).
		Return(singleResultNotFoundErr)

	netWorthCollection.(*mocks.Collection).
		On("FindOne", context.Background(), bson.M{"handle": "correctUserHandle"}).
		Return(singleResultFound)

	db.(*mocks.Database).
		On("Collection", "net_worth").
		Return(netWorthCollection)
	type fields struct {
		ctx context.Context
		db  mongo.Database
	}
	type args struct {
		ctx    context.Context
		handle string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			"erroneous user handle",
			fields{
				ctx: bctx,
				db:  db,
			},
			args{bctx, "errorUserHandle"},
			false,
			true,
		},
		{
			"not existing handle",
			fields{
				ctx: bctx,
				db:  db,
			},
			args{bctx, "notFoundHandle"},
			false,
			false,
		},

		{
			"existing handle",
			fields{
				ctx: bctx,
				db:  db,
			},
			args{bctx, "correctUserHandle"},
			true,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				ctx: tt.fields.ctx,
				db:  tt.fields.db,
			}
			got, err := r.exists(tt.args.ctx, tt.args.handle)
			if (err != nil) != tt.wantErr {
				t.Errorf("exists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("exists() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_getNetWorth(t *testing.T) {
	bctx := context.Background()
	var db mongo.Database
	var netWorthCollection mongo.Collection
	var singleResultErr mongo.SingleResult
	var singleResultNotFoundErr mongo.SingleResult
	var singleResultFound mongo.SingleResult

	db = &mocks.Database{}
	netWorthCollection = &mocks.Collection{}
	singleResultErr = &mocks.SingleResult{}
	singleResultNotFoundErr = &mocks.SingleResult{}
	singleResultFound = &mocks.SingleResult{}

	singleResultErr.(*mocks.SingleResult).
		On("Decode", mock.AnythingOfType("*net_worth.NetWorth")).
		Return(fmt.Errorf("mocked-error"))

	singleResultNotFoundErr.(*mocks.SingleResult).
		On("Decode", mock.AnythingOfType("*net_worth.NetWorth")).
		Return(mdb.ErrNoDocuments)

	singleResultFound.(*mocks.SingleResult).
		On("Decode", mock.AnythingOfType("*net_worth.NetWorth")).
		Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*NetWorth)
		arg.Handle = "correctUserHandle"
	})

	netWorthCollection.(*mocks.Collection).
		On("FindOne", context.Background(), bson.M{"handle": "errorUserHandle"}).
		Return(singleResultErr)

	netWorthCollection.(*mocks.Collection).
		On("FindOne", context.Background(), bson.M{"handle": "notFoundHandle"}).
		Return(singleResultNotFoundErr)

	netWorthCollection.(*mocks.Collection).
		On("FindOne", context.Background(), bson.M{"handle": "correctUserHandle"}).
		Return(singleResultFound)

	db.(*mocks.Database).
		On("Collection", "net_worth").
		Return(netWorthCollection)

	type fields struct {
		ctx context.Context
		db  mongo.Database
	}
	type args struct {
		ctx    context.Context
		handle string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *NetWorth
		wantErr bool
	}{
		{
			"erroneous user handle",
			fields{
				ctx: bctx,
				db:  db,
			},
			args{bctx, "errorUserHandle"},
			nil,
			true,
		},
		{
			"not existing handle",
			fields{
				ctx: bctx,
				db:  db,
			},
			args{bctx, "notFoundHandle"},
			nil,
			true,
		},

		{
			"existing handle",
			fields{
				ctx: bctx,
				db:  db,
			},
			args{bctx, "correctUserHandle"},
			&NetWorth{Handle: "correctUserHandle"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				ctx: tt.fields.ctx,
				db:  tt.fields.db,
			}
			got, err := r.getNetWorth(tt.args.ctx, tt.args.handle)
			if (err != nil) != tt.wantErr {
				t.Errorf("getNetWorth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNetWorth() got = %v, want %v", got, tt.want)
			}
		})
	}
}
