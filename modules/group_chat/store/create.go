package gchatstore

import (
	"context"
	"cs_chat_app_server/common"
	gchatmdl "cs_chat_app_server/modules/group_chat/model"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *mongoStore) Create(ctx context.Context, groupChatItem *gchatmdl.GroupChatItem) error {
	result, err := s.database.Collection(groupChatItem.CollectionName()).InsertOne(ctx, groupChatItem)
	if err != nil {
		log.Error().
			Err(err).
			Str("package", "gchatstore.Create").
			Msg("error while insert group chat item")
		return common.ErrInternal(err)
	}
	insertedId := result.InsertedID.(primitive.ObjectID).Hex()
	groupChatItem.Id = &insertedId
	return nil
}
