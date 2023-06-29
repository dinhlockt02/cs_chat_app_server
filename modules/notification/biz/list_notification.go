package notibiz

import (
	"context"
	notimodel "cs_chat_app_server/components/notification/model"
	notirepo "cs_chat_app_server/components/notification/repository"
	notistore "cs_chat_app_server/components/notification/store"
)

type listNotificationBiz struct {
	store notirepo.NotificationRepository
}

func NewListNotificationBiz(store notirepo.NotificationRepository) *listNotificationBiz {
	return &listNotificationBiz{store: store}
}

func (biz *listNotificationBiz) List(ctx context.Context, requesterId string) ([]notimodel.Notification, error) {
	filter := notistore.GetOwnerFilter(requesterId)
	return biz.store.List(ctx, filter)
}
