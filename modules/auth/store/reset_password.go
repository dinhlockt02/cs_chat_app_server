package authstore

import (
	"context"
	"cs_chat_app_server/common"
	authmodel "cs_chat_app_server/modules/auth/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *mongoStore) ResetPassword(ctx context.Context, filter map[string]interface{}, data *authmodel.ResetPasswordBody) error {
	update := bson.D{{"$set", data}}

	_, err := s.database.
		Collection(data.CollectionName()).
		UpdateOne(ctx, filter, update)
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
