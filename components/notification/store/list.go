package notistore

import (
	"context"
	"cs_chat_app_server/common"
	notimodel "cs_chat_app_server/components/notification/model"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *mongoStore) List(ctx context.Context, filter map[string]interface{}) ([]notimodel.Notification, error) {

	opts := options.Find().SetSort(bson.D{{"created_at", -1}})
	cursor, err := s.database.Collection(notimodel.Notification{}.CollectionName()).Find(ctx, filter, opts)
	if err != nil {
		log.Error().Err(err).Str("pakage", "notistore.Find").Send()
		return nil, common.ErrInternal(err)
	}
	var result []notimodel.Notification
	if err = cursor.All(ctx, &result); err != nil {
		log.Error().Err(err).Str("pakage", "notistore.Find").Send()
		return nil, common.ErrInternal(err)
	}

	return result, nil
}
