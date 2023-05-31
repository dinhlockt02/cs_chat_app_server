package grouprepo

import (
	"context"
	requestmdl "cs_chat_app_server/modules/request/model"
)

func (repo *groupRepository) CreateRequest(
	ctx context.Context,
	request *requestmdl.Request,
) error {
	return repo.requestStore.CreateRequest(ctx, request)
}
