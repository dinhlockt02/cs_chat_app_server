package gchatrepo

import (
	"context"
	"cs_chat_app_server/common"
	gchatmdl "cs_chat_app_server/modules/group_chat/model"
)

// TODO: Cache User Data
func (repo *GroupChatRepository) Create(ctx context.Context,
	groupChatItem *gchatmdl.GroupChatItem,
) error {
	err := repo.groupChatStore.Create(ctx, groupChatItem)
	if err != nil {
		return common.ErrInternal(err)
	}

	return nil
}
