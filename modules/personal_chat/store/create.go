package pchatstore

import (
	"context"
	"cs_chat_app_server/common"
	pchatmdl "cs_chat_app_server/modules/personal_chat/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *mongoStore) Create(ctx context.Context, personalChatItem *pchatmdl.PersonalChatItem) error {
	result, err := s.database.Collection(personalChatItem.CollectionName()).InsertOne(ctx, personalChatItem)
	if err != nil {
		return common.ErrInternal(err)
	}
	insertedId := result.InsertedID.(primitive.ObjectID).Hex()
	personalChatItem.Id = &insertedId
	return nil
}
