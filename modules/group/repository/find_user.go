package grouprepo

import (
	"context"
	groupmdl "cs_chat_app_server/modules/group/model"
)

func (repo *groupRepository) FindUser(ctx context.Context, filter map[string]interface{}) (*groupmdl.User, error) {
	return repo.groupStore.FindUser(ctx, filter)
}
