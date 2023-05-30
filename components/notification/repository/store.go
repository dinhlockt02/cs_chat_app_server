package notirepo

import (
	"context"
	notimodel "cs_chat_app_server/components/notification/model"
)

type NotificationStore interface {
	Create(ctx context.Context, data *notimodel.Notification) error
	FindDevice(ctx context.Context, filter map[string]interface{}) ([]notimodel.Device, error)
}
