package pchatstore

import (
	"context"
	"cs_chat_app_server/common"
	pchatmdl "cs_chat_app_server/modules/personal_chat/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *mongoStore) FindUser(ctx context.Context, filter map[string]interface{}) (*pchatmdl.User, error) {
	result := s.database.
		Collection(pchatmdl.User{}.CollectionName()).
		FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, common.ErrInternal(err)
	}
	var user = new(pchatmdl.User)
	if err := result.Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}
