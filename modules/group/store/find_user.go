package groupstore

import (
	"context"
	"cs_chat_app_server/common"
	groupmdl "cs_chat_app_server/modules/group/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *mongoStore) FindUser(ctx context.Context, filter map[string]interface{}) (*groupmdl.User, error) {
	var user groupmdl.User
	result := s.database.Collection(user.CollectionName()).FindOne(ctx, filter)

	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, common.ErrInternal(err)
	}

	if err := result.Decode(&user); err != nil {
		return nil, common.ErrInternal(err)
	}

	return &user, nil
}
