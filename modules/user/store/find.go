package userstore

import (
	"context"
	"cs_chat_app_server/common"
	usermodel "cs_chat_app_server/modules/user/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *mongoStore) Find(ctx context.Context, filter map[string]interface{}) (*usermodel.User, error) {
	result := s.database.
		Collection(usermodel.User{}.CollectionName()).
		FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, common.ErrInternal(err)
	}
	var user = new(usermodel.User)
	if err := result.Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}
