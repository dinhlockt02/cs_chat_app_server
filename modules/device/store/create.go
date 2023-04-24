package devicestore

import (
	"context"
	"cs_chat_app_server/common"
	devicemodel "cs_chat_app_server/modules/device/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *mongoStore) Create(ctx context.Context, data *devicemodel.Device) error {
	result, err := s.database.Collection(data.CollectionName()).InsertOne(ctx, data)
	if err != nil {
		return common.ErrInternal(err)
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	data.Id = &id
	return nil
}
