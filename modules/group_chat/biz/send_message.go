package gchatbiz

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/pubsub"
	gchatmdl "cs_chat_app_server/modules/group_chat/model"
	gchatrepo "cs_chat_app_server/modules/group_chat/repository"
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
		return common.ErrInternal(err)
	}

	biz.ps.Publish(ctx, common.TopicNewGroupMessageCreated, *item.Id)
	return nil
}
