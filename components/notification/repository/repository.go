package notirepo

import (
	"context"
	notimodel "cs_chat_app_server/components/notification/model"
	notistore "cs_chat_app_server/components/notification/store"
)

type notificationRepository struct {
	store notistore.NotificationStore
}

func NewNotificationRepository(store notistore.NotificationStore) *notificationRepository {
	return &notificationRepository{
		store: store,
	}
}

func (n *notificationRepository) List(ctx context.Context, filter map[string]interface{}) ([]notimodel.Notification, error) {
	return n.store.List(ctx, filter)
}

func (n *notificationRepository) Delete(ctx context.Context, filter map[string]interface{}) error {
	return n.store.Delete(ctx, filter)
}
