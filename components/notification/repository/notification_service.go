package notirepo

import (
	"context"
	notimodel "cs_chat_app_server/components/notification/model"
)

type NotificationService interface {
	Push(ctx context.Context, token []string, notification *notimodel.Notification) error
}
