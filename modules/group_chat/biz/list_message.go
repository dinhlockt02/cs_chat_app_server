package gchatbiz

import (
	"context"
	"cs_chat_app_server/common"
	gchatmdl "cs_chat_app_server/modules/group_chat/model"
	gchatrepo "cs_chat_app_server/modules/group_chat/repository"
)

type listMessageBiz struct {
	groupChatRepo gchatrepo.Repository
}

func NewListMessageBiz(
	groupChatRepo gchatrepo.Repository,
) *listMessageBiz {
	return &listMessageBiz{
		groupChatRepo: groupChatRepo,
	}
}

func (biz *listMessageBiz) List(
	ctx context.Context,
	requesterId string,
	filter map[string]interface{},
	paging gchatmdl.Paging,
) ([]gchatmdl.GroupChatItem, error) {
	list, err := biz.groupChatRepo.List(ctx, filter, paging)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	t := true
	f := false
	for i := range list {
		if list[i].SenderId == requesterId {
			list[i].IsMe = &t
		} else {
			list[i].IsMe = &f
		}
	}
	return list, nil
}
