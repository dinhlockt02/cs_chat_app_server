package gchatstore

import (
	"context"
	"cs_chat_app_server/common"
	gchatmdl "cs_chat_app_server/modules/group_chat/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *mongoStore) Create(ctx context.Context, groupChatItem *gchatmdl.GroupChatItem) error {
	result, err := s.database.Collection(groupChatItem.CollectionName()).InsertOne(ctx, groupChatItem)
	if err != nil {
		return common.ErrInternal(err)
	}
	insertedId := result.InsertedID.(primitive.ObjectID).Hex()
	groupChatItem.Id = &insertedId
	return nil
}
