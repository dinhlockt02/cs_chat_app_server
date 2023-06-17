package devicestore

import (
	"context"
	devicemodel "cs_chat_app_server/modules/device/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store interface {
	Create(ctx context.Context, data *devicemodel.Device) error
	Update(ctx context.Context, filter map[string]interface{}, data *devicemodel.UpdateDevice) error
	Delete(ctx context.Context, filter map[string]interface{}) error
	Get(ctx context.Context, filter map[string]interface{}) ([]*devicemodel.GetDeviceDto, error)
}

type mongoStore struct {
	database *mongo.Database
}

func NewMongoStore(database *mongo.Database) *mongoStore {
	return &mongoStore{database: database}
}
