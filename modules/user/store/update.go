package userstore

import (
	"context"
	"cs_chat_app_server/common"
	usermodel "cs_chat_app_server/modules/user/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *mongoStore) Update(ctx context.Context, filter map[string]interface{}, updatedUser *usermodel.UpdateUser) error {
	update := bson.D{{"$set", updatedUser}}

	_, err := s.database.
		Collection(updatedUser.CollectionName()).
		UpdateOne(ctx, filter, update)
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
