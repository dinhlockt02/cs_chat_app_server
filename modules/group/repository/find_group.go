package grouprepo

import (
	"context"
	groupmdl "cs_chat_app_server/modules/group/model"
)

// FindGroup returns the group using filter
func (repo *groupRepository) FindGroup(
	ctx context.Context,
	filter map[string]interface{},
) (*groupmdl.Group, error) {
	return repo.groupStore.FindGroup(ctx, filter)
}
