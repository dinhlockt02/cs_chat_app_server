package notirepo

import (
	"context"
	"cs_chat_app_server/common"
	notimodel "cs_chat_app_server/components/notification/model"
	notistore "cs_chat_app_server/components/notification/store"
	"github.com/rs/zerolog/log"
)

type notificationServiceRepository struct {
	service NotificationService
	store   notistore.NotificationStore
}

func NewNotificationServiceRepository(service NotificationService, store notistore.NotificationStore) *notificationServiceRepository {
	return &notificationServiceRepository{
		service: service,
		store:   store,
	}
}

func (repo *notificationServiceRepository) CreateAcceptFriendNotification(
	ctx context.Context,
	owner string,
	subject *notimodel.NotificationObject,
	indirect *notimodel.NotificationObject,
) error {
	noti := notimodel.
		NewNotificationBuilder(notimodel.AcceptRequest, owner).
		SetSubject(subject).
		SetIndirect(indirect).
		Build()

	return repo.createNotification(ctx, noti)
}

func (repo *notificationServiceRepository) CreateReceiveFriendRequestNotification(
	ctx context.Context,
	owner string,
	subject *notimodel.NotificationObject,
	prep *notimodel.NotificationObject,
) error {
	noti := notimodel.
		NewNotificationBuilder(notimodel.ReceiveFriendRequest, owner).
		SetSubject(subject).
		SetPrep(prep).
		Build()

	return repo.createNotification(ctx, noti)
}

func (repo *notificationServiceRepository) createNotification(ctx context.Context, noti *notimodel.Notification) error {
	err := repo.store.Create(ctx, noti)
	if err != nil {
		return common.ErrInternal(err)
	}

	go func() {
		devices, e := repo.store.FindDevice(context.Background(), map[string]interface{}{
			"user_id": noti.Owner,
		})
		if e != nil {
			log.Err(e)
			return
		}

		tokens := make([]string, 0, len(devices))

		for i := range devices {
			if token := devices[i].PushNotificationToken; token != "" {
				tokens = append(tokens, token)
			}
		}

		if len(tokens) == 0 {
			return
		}

		e = repo.service.Push(context.Background(), tokens, noti)
		if e != nil {
			log.Err(e)
		}
	}()

	return nil
}
