package authstore

import (
	"context"
	"cs_chat_app_server/common"
	authmodel "cs_chat_app_server/modules/auth/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *mongoStore) UpdateEmailVerified(ctx context.Context, filter map[string]interface{}) error {

	var updateEmailVerifiedUser authmodel.EmailVerifiedUser
	updateEmailVerifiedUser.Process()

	update := bson.D{{"$set", updateEmailVerifiedUser}}

	_, err := s.database.
		Collection(updateEmailVerifiedUser.CollectionName()).
		UpdateOne(ctx, filter, update)
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
