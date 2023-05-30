package friendrepo

import (
	"context"
)

// DeleteRequest delete requests that matches filter
func (repo *friendRepository) DeleteRequest(
	ctx context.Context,
	filter map[string]interface{},
) error {
	return repo.requestStore.DeleteRequest(ctx, filter)
}
