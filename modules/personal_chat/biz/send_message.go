package pchatbiz

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/socket"
	pchatmdl "cs_chat_app_server/modules/personal_chat/model"
)

type SendMessagePersonalChatStore interface {
	Create(ctx context.Context, personalChatItem *pchatmdl.PersonalChatItem) error
}

type sendMessageBiz struct {
	personalChatStore SendMessagePersonalChatStore
	skt               socket.Socket
}

func NewSendMessageBiz(
	personalChatStore SendMessagePersonalChatStore,
	skt socket.Socket,
) *sendMessageBiz {
	return &sendMessageBiz{
		personalChatStore: personalChatStore,
		skt:               skt,
	}
}

func (biz *sendMessageBiz) Send(ctx context.Context, item *pchatmdl.PersonalChatItem) error {
	if err := item.Process(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	err := biz.personalChatStore.Create(ctx, item)
	if err != nil {
		return common.ErrInternal(err)
	}
	err = biz.skt.Send(item.Receiver, item)
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
