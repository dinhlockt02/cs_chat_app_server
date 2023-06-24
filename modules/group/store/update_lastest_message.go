package groupstore

import (
	"context"
	"cs_chat_app_server/common"
	groupmdl "cs_chat_app_server/modules/group/model"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *mongoStore) UpdateLatestMessage(
	ctx context.Context,
	filter map[string]interface{},
	latestMessage *groupmdl.GroupMessage) error {
	updateData := bson.D{{
		"$set", map[string]interface{}{
			"latest_message": latestMessage,
		},
	}}
	rs, err := s.database.
		Collection(groupmdl.Group{}.CollectionName()).
		UpdateOne(ctx, filter, updateData)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		return common.ErrInternal(err)
	}
	log.Debug().Msgf("UpdateLatestMessage: %v matched, %v modified, %v Upserted", rs.ModifiedCount, rs.UpsertedCount, rs.UpsertedCount)
	return nil
}
