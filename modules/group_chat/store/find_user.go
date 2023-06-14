package gchatstore

import (
	"context"
	"cs_chat_app_server/common"
	gchatmdl "cs_chat_app_server/modules/group_chat/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *mongoStore) FindUser(ctx context.Context, filter map[string]interface{}) (*gchatmdl.User, error) {
	result := s.database.
		Collection(gchatmdl.User{}.CollectionName()).
		FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, common.ErrInternal(err)
	}
	var user = new(gchatmdl.User)
	if err := result.Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}
