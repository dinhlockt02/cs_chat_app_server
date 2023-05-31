package grouprepo

import (
	"context"
)

// DeleteRequest delete requests that matches filter
func (repo *groupRepository) DeleteRequest(
	ctx context.Context,
	filter map[string]interface{},
) error {
	return repo.requestStore.DeleteRequest(ctx, filter)
}
