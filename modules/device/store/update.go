package devicestore

import (
	"context"
	devicemodel "cs_chat_app_server/modules/device/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *mongoStore) Update(ctx context.Context, filter map[string]interface{}, data *devicemodel.UpdateDevice) error {
	update := bson.D{{"$set", data}}

	_, err := s.database.Collection(data.CollectionName()).UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
