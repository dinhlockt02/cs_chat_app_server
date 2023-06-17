package devicebiz

import (
	"context"
	devicemodel "cs_chat_app_server/modules/device/model"
	devicestore "cs_chat_app_server/modules/device/store"
)

type getDevicesBiz struct {
	store devicestore.Store
}

func NewGetDevicesBiz(store devicestore.Store) *getDevicesBiz {
	return &getDevicesBiz{store: store}
}

func (biz *getDevicesBiz) Get(ctx context.Context, filter map[string]interface{}) ([]*devicemodel.GetDeviceDto, error) {
	return biz.store.Get(ctx, filter)
}
