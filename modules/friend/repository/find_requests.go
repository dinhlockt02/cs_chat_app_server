package friendrepo

import (
	"context"
	requestmdl "cs_chat_app_server/modules/request/model"
)

// FindRequests returns the friend request between sender and receiver
// If the request does not exist, it returns nil, nil
func (repo *friendRepository) FindRequests(
	ctx context.Context,
	filter map[string]interface{},
) ([]requestmdl.Request, error) {
	return repo.requestStore.FindRequests(ctx, filter)
}
