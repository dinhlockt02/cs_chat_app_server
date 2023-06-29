package notistore

import (
	"context"
	"cs_chat_app_server/common"
	notimodel "cs_chat_app_server/components/notification/model"
	"github.com/rs/zerolog/log"
)

func (s *mongoStore) Delete(ctx context.Context, filter map[string]interface{}) error {

	_, err := s.database.Collection(notimodel.Notification{}.CollectionName()).DeleteMany(ctx, filter)
	if err != nil {
		log.Error().Err(err).Str("pakage", "notistore.Find").Send()
		return common.ErrInternal(err)
	}
	return nil
}
