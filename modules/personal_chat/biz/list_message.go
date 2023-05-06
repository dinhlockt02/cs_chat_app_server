package pchatbiz

import (
	"context"
	"cs_chat_app_server/common"
	pchatmdl "cs_chat_app_server/modules/personal_chat/model"
)

type ListMessagePersonalRepo interface {
	List(ctx context.Context,
		filter map[string]interface{},
		paging pchatmdl.Paging,
	) ([]pchatmdl.PersonalChatItem, error)
}

type listMessageBiz struct {
	personalChatRepo ListMessagePersonalRepo
}

func NewListMessageBiz(
	personalChatRepo ListMessagePersonalRepo,
) *listMessageBiz {
	return &listMessageBiz{
		personalChatRepo: personalChatRepo,
	}
}

func (biz *listMessageBiz) List(ctx context.Context,
	filter map[string]interface{},
	paging pchatmdl.Paging,
) ([]pchatmdl.PersonalChatItem, error) {
	list, err := biz.personalChatRepo.List(ctx, filter, paging)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return list, nil
}
