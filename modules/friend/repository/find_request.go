package friendrepo

import (
	"context"
	"cs_chat_app_server/common"
	requestmdl "cs_chat_app_server/modules/request/model"
	requeststore "cs_chat_app_server/modules/request/store"
)

// FindRequest returns the friend request between sender and receiver
// If the request does not exist, it returns nil, nil
func (repo *friendRepository) FindRequest(
	ctx context.Context,
	sender string,
	receiver string,
) (*requestmdl.Request, error) {

	senderFilter := requeststore.GetRequestSenderIdFilter(sender)
	receiverFilter := requeststore.GetRequestReceiverIdFilter(receiver)
	filter := common.GetAndFilter(senderFilter, receiverFilter)

	existedRequest, err := repo.requestStore.FindRequest(ctx, filter)
	if err != nil {
		return nil, err
	}
	return existedRequest, nil
}
