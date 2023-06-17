package grouprepo

import (
	"context"
	groupmdl "cs_chat_app_server/modules/group/model"
)

func (repo *groupRepository) FindUsers(ctx context.Context, filter map[string]interface{}) ([]groupmdl.User, error) {
	return repo.groupStore.FindUsers(ctx, filter)
}
