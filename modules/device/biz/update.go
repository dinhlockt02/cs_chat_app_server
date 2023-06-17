package devicebiz

import (
	"context"
	devicemodel "cs_chat_app_server/modules/device/model"
	devicestore "cs_chat_app_server/modules/device/store"
)

type createDeviceBiz struct {
	store devicestore.Store
}

func NewUpdateDeviceBiz(store devicestore.Store) *createDeviceBiz {
	return &createDeviceBiz{store: store}
}

func (biz *createDeviceBiz) Update(ctx context.Context, filter map[string]interface{}, data *devicemodel.UpdateDevice) error {

	if err := data.Process(); err != nil {
		return err
	}

	err := biz.store.Update(ctx, filter, data)
	if err != nil {
		return err
	}

	return nil
}
