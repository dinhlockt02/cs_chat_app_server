package gchatbiz

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/pubsub"
	gchatmdl "cs_chat_app_server/modules/group_chat/model"
	gchatrepo "cs_chat_app_server/modules/group_chat/repository"
	"github.com/rs/zerolog/log"
)

type sendMessageBiz struct {
	groupChatRepo gchatrepo.Repository
	ps            pubsub.PubSub
}

func NewSendMessageBiz(
	groupChatRepo gchatrepo.Repository,
	ps pubsub.PubSub,
) *sendMessageBiz {
	return &sendMessageBiz{
		groupChatRepo: groupChatRepo,
		ps:            ps,
	}
}

func (biz *sendMessageBiz) Send(ctx context.Context, item *gchatmdl.GroupChatItem) error {
	if err := item.Process(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	err := biz.groupChatRepo.Create(ctx, item)
	if err != nil {
		log.Error().
			Err(err).
			Str("package", "gchatbiz.Send").
			Msg("error while biz.groupChatRepo.Create")
		return common.ErrInternal(err)
	}

	log.Debug().
		Err(err).
		Str("package", "gchatbiz.Send").
		Msg("publishing TopicNewGroupMessageCreated with payload: " + *item.Id)
	biz.ps.Publish(ctx, common.TopicNewGroupMessageCreated, *item.Id)
	return nil
}
