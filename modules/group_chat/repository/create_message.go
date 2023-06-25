package gchatrepo

import (
	"context"
	"cs_chat_app_server/common"
	gchatmdl "cs_chat_app_server/modules/group_chat/model"
	"github.com/rs/zerolog/log"
)

// TODO: Cache User Data
func (repo *GroupChatRepository) Create(ctx context.Context,
	groupChatItem *gchatmdl.GroupChatItem,
) error {
	err := repo.groupChatStore.Create(ctx, groupChatItem)
	if err != nil {
		log.Error().
			Err(err).
			Str("package", "gchatrepo.Create").
			Msg("error while repo.groupChatStore.Create")
		return common.ErrInternal(err)
	}

	return nil
}
