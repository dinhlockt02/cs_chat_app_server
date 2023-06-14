package gchatrepo

import (
	"context"
	gchatmdl "cs_chat_app_server/modules/group_chat/model"
	gchatstore "cs_chat_app_server/modules/group_chat/store"
)

type Repository interface {
	List(
		ctx context.Context,
		filter map[string]interface{},
		paging gchatmdl.Paging,
	) ([]gchatmdl.GroupChatItem, error)
	Create(ctx context.Context,
		groupChatItem *gchatmdl.GroupChatItem,
	) error
}

type GroupChatRepository struct {
	groupChatStore gchatstore.Store
}

func NewGroupChatRepository(groupChatStore gchatstore.Store) *GroupChatRepository {
	return &GroupChatRepository{groupChatStore: groupChatStore}
}
