package requeststore

import (
	"context"
	requestmdl "cs_chat_app_server/modules/request/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store interface {
	FindRequests(ctx context.Context, filter map[string]interface{}) ([]requestmdl.Request, error)
	FindRequest(ctx context.Context, filter map[string]interface{}) (*requestmdl.Request, error)
	DeleteRequest(ctx context.Context, filter map[string]interface{}) error
	CreateRequest(ctx context.Context, request *requestmdl.Request) error
}

type mongoStore struct {
	database *mongo.Database
}

func NewMongoStore(database *mongo.Database) *mongoStore {
	return &mongoStore{database: database}
}
