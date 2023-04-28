package authstore

import (
	"context"
	"cs_chat_app_server/common"
	authmodel "cs_chat_app_server/modules/auth/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *mongoStore) Find(ctx context.Context, filter map[string]interface{}) (*authmodel.User, error) {

	var findUser authmodel.User
	result := s.database.Collection(findUser.CollectionName()).FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, common.ErrInternal(err)
	}
	err := result.Decode(&findUser)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return &findUser, nil
}
