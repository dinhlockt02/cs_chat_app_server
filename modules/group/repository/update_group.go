package grouprepo

import (
	"context"
	groupmdl "cs_chat_app_server/modules/group/model"
)

func (repo *groupRepository) UpdateGroup(
	ctx context.Context,
	filter map[string]interface{},
	updatedGroup *groupmdl.UpdateGroup,
) error {
	return repo.groupStore.UpdateGroup(ctx, filter, updatedGroup)
}
