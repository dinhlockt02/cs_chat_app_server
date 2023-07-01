package grouprepo

import (
	"context"
)

func (repo *groupRepository) DeleteGroups(ctx context.Context, filter map[string]interface{}) error {
	return repo.groupStore.DeleteGroups(ctx, filter)
}
