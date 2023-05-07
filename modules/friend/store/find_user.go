package friendstore

import (
	"context"
	"cs_chat_app_server/common"
	friendmodel "cs_chat_app_server/modules/friend/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// FindUser is a method for finding a User
// which is filtered by a map[string] interface
func (s *mongoStore) FindUser(ctx context.Context, filter map[string]interface{}) (*friendmodel.User, error) {
	var user friendmodel.User
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
