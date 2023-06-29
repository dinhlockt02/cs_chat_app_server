package notibiz

import (
	"context"
	notirepo "cs_chat_app_server/components/notification/repository"
	notistore "cs_chat_app_server/components/notification/store"
)

type deleteAllNotificationBiz struct {
	store notirepo.NotificationRepository
}

func NewDeleteAllNotificationBiz(store notirepo.NotificationRepository) *deleteAllNotificationBiz {
	return &deleteAllNotificationBiz{store: store}
}

func (biz *deleteAllNotificationBiz) DeleteAll(ctx context.Context, requesterId string) error {
	filter := notistore.GetOwnerFilter(requesterId)
	return biz.store.Delete(ctx, filter)
}
