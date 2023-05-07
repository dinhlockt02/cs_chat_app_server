package pchatbiz

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/socket"
	pchatmdl "cs_chat_app_server/modules/personal_chat/model"
)

type SendMessagePersonalChatRepo interface {
	Create(ctx context.Context,
		personalChatItem *pchatmdl.PersonalChatItem,
	) error
}

type sendMessageBiz struct {
	personalChatRepo SendMessagePersonalChatRepo
	skt              socket.Socket
}

func NewSendMessageBiz(
	personalChatStore SendMessagePersonalChatRepo,
	skt socket.Socket,
) *sendMessageBiz {
	return &sendMessageBiz{
		personalChatRepo: personalChatStore,
		skt:              skt,
	}
}

func (biz *sendMessageBiz) Send(ctx context.Context, item *pchatmdl.PersonalChatItem) error {
	if err := item.Process(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	err := biz.personalChatRepo.Create(ctx, item)
	if err != nil {
		return common.ErrInternal(err)
	}

	t := true
	f := false

	if item.SenderId != item.ReceiverId {
		item.IsMe = &t
		err = biz.skt.Send(item.SenderId, item)
	}
	item.IsMe = &f
	err = biz.skt.Send(item.ReceiverId, item)
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
