package friendrepo

import (
	"context"
	friendmodel "cs_chat_app_server/modules/friend/model"
)

// UpdateUser updates user that matches filter
func (repo *friendRepository) UpdateUser(
	ctx context.Context,
	filter map[string]interface{},
	user *friendmodel.User,
) error {
	return repo.friendstore.UpdateUser(ctx, filter, user)
}
