package authstore

import (
	"context"
	"cs_chat_app_server/common"
	authmodel "cs_chat_app_server/modules/auth/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *mongoStore) Update(ctx context.Context, filter map[string]interface{}, passwordUser *authmodel.UpdatePasswordUser) error {
	update := bson.D{{"$set", passwordUser}}

	_, err := s.database.
		Collection(passwordUser.CollectionName()).
		UpdateOne(ctx, filter, update)
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
