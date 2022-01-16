package mongo

import (
	"context"
	"github.com/avinashb98/myfin/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database interface {
	Collection(name string) Collection
	Client() Client
}

type Collection interface {
	FindOne(context.Context, interface{}) SingleResult
	InsertOne(context.Context, interface{}) (interface{}, error)
	//DeleteOne(ctx context.Context, filter interface{}) (int64, error)
}

type SingleResult interface {
	Decode(v interface{}) error
}

type Client interface {
	Database(string) Database
	Connect() error
	//StartSession() (mongo.Session, error)
}

type mongoClient struct {
	cl *mongo.Client
}
type mongoDatabase struct {
	db *mongo.Database
}
type mongoCollection struct {
	coll *mongo.Collection
}

type mongoSingleResult struct {
	sr *mongo.SingleResult
}

type mongoSession struct {
	mongo.Session
}

func NewClient(cnf *config.MongoConfig) (Client, error) {
	c, err := mongo.NewClient(options.Client().ApplyURI(cnf.URI))
	return &mongoClient{cl: c}, err
}

func NewDatabase(cnf *config.MongoConfig, client Client) Database {
	return client.Database(cnf.Database)
}

func (mc *mongoClient) Database(dbName string) Database {
	db := mc.cl.Database(dbName)
	return &mongoDatabase{db: db}
}

//func (mc *mongoClient) StartSession() (mongo.Session, error) {
//	session, err := mc.cl.StartSession()
//	return &mongoSession{session}, err
//}

func (mc *mongoClient) Connect() error {
	return mc.cl.Connect(nil)
}

func (md *mongoDatabase) Collection(colName string) Collection {
	collection := md.db.Collection(colName)
	return &mongoCollection{coll: collection}
}

func (md *mongoDatabase) Client() Client {
	client := md.db.Client()
	return &mongoClient{cl: client}
}

func (mc *mongoCollection) FindOne(ctx context.Context, filter interface{}) SingleResult {
	singleResult := mc.coll.FindOne(ctx, filter)
	return &mongoSingleResult{sr: singleResult}
}

func (mc *mongoCollection) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	id, err := mc.coll.InsertOne(ctx, document)
	return id.InsertedID, err
}

//func (mc *mongoCollection) DeleteOne(ctx context.Context, filter interface{}) (int64, error) {
//	count, err := mc.coll.DeleteOne(ctx, filter)
//	return count.DeletedCount, err
//}

func (sr *mongoSingleResult) Decode(v interface{}) error {
	return sr.sr.Decode(v)
}
