package grouprepo

import (
	"context"
	groupmdl "cs_chat_app_server/modules/group/model"
)

func (repo *groupRepository) List(ctx context.Context, groupFilter map[string]interface{}) ([]groupmdl.Group, error) {
	return repo.groupStore.List(ctx, groupFilter)
}
