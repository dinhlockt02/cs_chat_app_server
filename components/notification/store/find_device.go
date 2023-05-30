package notistore

import (
	"context"
	"cs_chat_app_server/common"
	notimodel "cs_chat_app_server/components/notification/model"
)

func (s *mongoStore) FindDevice(ctx context.Context, filter map[string]interface{}) ([]notimodel.Device, error) {
	cursor, err := s.database.Collection(notimodel.Device{}.CollectionName()).
		Find(ctx, filter)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	var devices []notimodel.Device
	if err = cursor.All(ctx, &devices); err != nil {
		return nil, common.ErrInternal(err)
	}
	return devices, nil
}
