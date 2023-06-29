package notibiz

import (
	"context"
	"cs_chat_app_server/common"
	notirepo "cs_chat_app_server/components/notification/repository"
	notistore "cs_chat_app_server/components/notification/store"
	"github.com/rs/zerolog/log"
)

type deleteByIdNotificationBiz struct {
	store notirepo.NotificationRepository
}

func NewDeleteByIdNotificationBiz(store notirepo.NotificationRepository) *deleteByIdNotificationBiz {
	return &deleteByIdNotificationBiz{store: store}
}

func (biz *deleteByIdNotificationBiz) DeleteById(ctx context.Context, requesterId, id string) error {
	filter, err := common.GetIdFilter(id)
	if err != nil {
		log.Debug().Err(err).Str("package", "notibiz.DeleteById").Send()
		return err
	}
	filter = common.GetAndFilter(filter, notistore.GetOwnerFilter(requesterId))
	return biz.store.Delete(ctx, filter)
}
