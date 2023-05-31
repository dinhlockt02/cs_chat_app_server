package grouprepo

import (
	"context"
	groupmdl "cs_chat_app_server/modules/group/model"
)

func (repo *groupRepository) CreateGroup(ctx context.Context, group *groupmdl.Group) error {
	return repo.groupStore.Create(ctx, group)
}
