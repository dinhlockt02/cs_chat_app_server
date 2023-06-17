package devicestore

import (
	"context"
	"cs_chat_app_server/common"
	devicemodel "cs_chat_app_server/modules/device/model"
)

func (s *mongoStore) Get(ctx context.Context, filter map[string]interface{}) ([]*devicemodel.GetDeviceDto, error) {
	var devices []*devicemodel.GetDeviceDto
	cursor, err := s.database.Collection(devicemodel.GetDeviceDto{}.CollectionName()).Find(ctx, filter)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	if err = cursor.All(ctx, &devices); err != nil {
		return nil, common.ErrInvalidRequest(err)
	}
	return devices, nil
}
