package notistore

import (
	"context"
	"cs_chat_app_server/common"
	notimodel "cs_chat_app_server/components/notification/model"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *mongoStore) UpdateNotifications(
	ctx context.Context,
	filter map[string]interface{},
	data *notimodel.UpdateNotification,
) error {
	updateData := bson.D{{
		"$set", data,
	}}
	rs, err := s.database.Collection(data.CollectionName()).UpdateMany(ctx, filter, updateData)
	if err != nil {
		return common.ErrInternal(err)
	}
	log.Debug().Msgf("UpdateNotifications: %v matched, %v modified, %v Upserted", rs.ModifiedCount, rs.UpsertedCount, rs.UpsertedCount)
	return nil
}
