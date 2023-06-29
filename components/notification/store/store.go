package notistore

import (
	"context"
	notimodel "cs_chat_app_server/components/notification/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationStore interface {
	Create(ctx context.Context, data *notimodel.Notification) error
	FindDevice(ctx context.Context, filter map[string]interface{}) ([]notimodel.Device, error)
	List(ctx context.Context, filter map[string]interface{}) ([]notimodel.Notification, error)
	Delete(ctx context.Context, filter map[string]interface{}) error
}

type mongoStore struct {
	database *mongo.Database
}

func NewMongoStore(database *mongo.Database) *mongoStore {
	return &mongoStore{database: database}
}
