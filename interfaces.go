package mongokit

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

var _ Client = &mongo.Client{}
var _ Database = &mongo.Database{}
var _ Collection = &mongo.Collection{}

// Database is an interface for mongo.Database
type Database interface {
	Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, error)
	Client() *mongo.Client
	Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection
	CreateCollection(ctx context.Context, name string, opts ...*options.CreateCollectionOptions) error
	CreateView(ctx context.Context, viewName, viewOn string, pipeline interface{}, opts ...*options.CreateViewOptions) error
	Drop(ctx context.Context) error
	ListCollectionNames(ctx context.Context, filter interface{}, opts ...*options.ListCollectionsOptions) ([]string, error)
	ListCollectionSpecifications(ctx context.Context, filter interface{}, opts ...*options.ListCollectionsOptions) ([]*mongo.CollectionSpecification, error)
	ListCollections(ctx context.Context, filter interface{}, opts ...*options.ListCollectionsOptions) (*mongo.Cursor, error)
	Name() string
	ReadConcern() *readconcern.ReadConcern
	ReadPreference() *readpref.ReadPref
	RunCommand(ctx context.Context, runCommand interface{}, opts ...*options.RunCmdOptions) *mongo.SingleResult
	RunCommandCursor(ctx context.Context, runCommand interface{}, opts ...*options.RunCmdOptions) (*mongo.Cursor, error)
	Watch(ctx context.Context, pipeline interface{}, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error)
	WriteConcern() *writeconcern.WriteConcern
}

// Client is an interface for mongo.Client
type Client interface {
	Connect(ctx context.Context) error
	Database(name string, opts ...*options.DatabaseOptions) *mongo.Database
	Disconnect(ctx context.Context) error
	ListDatabaseNames(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions) ([]string, error)
	ListDatabases(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions) (mongo.ListDatabasesResult, error)
	NumberSessionsInProgress() int
	Ping(ctx context.Context, rp *readpref.ReadPref) error
	StartSession(opts ...*options.SessionOptions) (mongo.Session, error)
	Timeout() *time.Duration
	UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error
	UseSessionWithOptions(ctx context.Context, opts *options.SessionOptions, fn func(mongo.SessionContext) error) error
	Watch(ctx context.Context, pipeline interface{}, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error)
}

// Collection is an interface for mongo.Collection
type Collection interface {
	Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, error)
	BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error)
	Clone(opts ...*options.CollectionOptions) (*mongo.Collection, error)
	CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error)
	Database() *mongo.Database
	DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	Distinct(ctx context.Context, fieldName string, filter interface{}, opts ...*options.DistinctOptions) ([]interface{}, error)
	Drop(ctx context.Context) error
	EstimatedDocumentCount(ctx context.Context, opts ...*options.EstimatedDocumentCountOptions) (int64, error)
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (cur *mongo.Cursor, err error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	FindOneAndDelete(ctx context.Context, filter interface{}, opts ...*options.FindOneAndDeleteOptions) *mongo.SingleResult
	FindOneAndReplace(ctx context.Context, filter interface{}, replacement interface{}, opts ...*options.FindOneAndReplaceOptions) *mongo.SingleResult
	FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult
	Indexes() mongo.IndexView
	InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	Name() string
	ReplaceOne(ctx context.Context, filter interface{}, replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error)
	SearchIndexes() mongo.SearchIndexView
	UpdateByID(ctx context.Context, id interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	Watch(ctx context.Context, pipeline interface{}, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error)
}
