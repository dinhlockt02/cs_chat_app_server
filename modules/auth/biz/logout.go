package authbiz

import (
	"context"
)

type LogoutDeviceStore interface {
	Delete(ctx context.Context, filter map[string]interface{}) error
}

type logoutBiz struct {
	deviceStore LogoutDeviceStore
}

func NewLogoutBiz(
	deviceStore LogoutDeviceStore,
) *logoutBiz {
	return &logoutBiz{
		deviceStore: deviceStore,
	}
}

func (biz *logoutBiz) Logout(ctx context.Context, filter map[string]interface{}) error {
	err := biz.deviceStore.Delete(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
