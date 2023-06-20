package pchatbiz

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/pubsub"
	pchatmdl "cs_chat_app_server/modules/personal_chat/model"
	"github.com/rs/zerolog/log"
)

type SendMessagePersonalChatRepo interface {
	Create(ctx context.Context,
		personalChatItem *pchatmdl.PersonalChatItem,
	) error
}

type sendMessageBiz struct {
	personalChatRepo SendMessagePersonalChatRepo
	ps               pubsub.PubSub
}

func NewSendMessageBiz(
	personalChatStore SendMessagePersonalChatRepo,
	ps pubsub.PubSub,
) *sendMessageBiz {
	return &sendMessageBiz{
		personalChatRepo: personalChatStore,
		ps:               ps,
	}
}

func (biz *sendMessageBiz) Send(ctx context.Context, item *pchatmdl.PersonalChatItem) error {
	if err := item.Process(); err != nil {
		log.Debug().Err(err).Msg("biz.Send: invalid request")
		return common.ErrInvalidRequest(err)
	}

	err := biz.personalChatRepo.Create(ctx, item)
	if err != nil {
		log.Debug().Err(err).Msg("biz.Send: create failed")
		return common.ErrInternal(err)
	}
	biz.ps.Publish(ctx, common.TopicNewPersonalMessageCreated, *item.Id)
	return nil
}
