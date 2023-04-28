package devicestore

import (
	"context"
	"cs_chat_app_server/common"
	devicemodel "cs_chat_app_server/modules/device/model"
)

func (s *mongoStore) Delete(ctx context.Context, filter map[string]interface{}) error {
	_, err := s.database.Collection(devicemodel.Device{}.CollectionName()).DeleteMany(ctx, filter)
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
