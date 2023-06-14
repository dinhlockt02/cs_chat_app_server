package pchatstore

import (
	"context"
	"cs_chat_app_server/common"
	pchatmdl "cs_chat_app_server/modules/personal_chat/model"

	"go.mongodb.org/mongo-driver/mongo"
)

func (s *mongoStore) FindMessage(ctx context.Context, filter map[string]interface{}) (*pchatmdl.PersonalChatItem, error) {
	var item pchatmdl.PersonalChatItem
	result := s.database.
		Collection(item.CollectionName()).
		FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, common.ErrInternal(err)
	}
	if err := result.Decode(&item); err != nil {
		return nil, err
	}
	return &item, nil
}
