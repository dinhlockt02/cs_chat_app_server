package authmiddleware

import (
	"context"
	"cs_chat_app_server/common"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoStore struct {
	database *mongo.Database
}

func NewMongoStore(database *mongo.Database) *mongoStore {
	return &mongoStore{database: database}
}

func (s *mongoStore) FindOne(ctx context.Context, filter map[string]interface{}) (*Device, error) {
	var device Device
	result := s.database.Collection(device.CollectionName()).FindOne(ctx, filter)

	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, common.ErrInternal(err)
	}
	if err := result.Decode(&device); err != nil {
		log.Debug().Msg(err.Error())

		return nil, common.ErrInternal(err)
	}
	return &device, nil
}
