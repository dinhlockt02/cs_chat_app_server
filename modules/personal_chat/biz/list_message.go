package pchatbiz

import (
	"context"
	"cs_chat_app_server/common"
	pchatmdl "cs_chat_app_server/modules/personal_chat/model"
)

type ListMessagePersonalChatStore interface {
	List(
		ctx context.Context,
		filter map[string]interface{},
		paging pchatmdl.Paging,
	) ([]pchatmdl.PersonalChatItem, error)
}

type listMessageBiz struct {
	personalChatStore ListMessagePersonalChatStore
}

func NewListMessageBiz(
	personalChatStore ListMessagePersonalChatStore,
) *listMessageBiz {
	return &listMessageBiz{
		personalChatStore: personalChatStore,
	}
}

func (biz *listMessageBiz) List(ctx context.Context,
	filter map[string]interface{},
	paging pchatmdl.Paging,
) ([]pchatmdl.PersonalChatItem, error) {
	list, err := biz.personalChatStore.List(ctx, filter, paging)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return list, nil
}
