package groupbiz

import (
	"context"
	"cs_chat_app_server/common"
	notirepo "cs_chat_app_server/components/notification/repository"
	groupmdl "cs_chat_app_server/modules/group/model"
	grouprepo "cs_chat_app_server/modules/group/repository"
	requestmdl "cs_chat_app_server/modules/request/model"
	requeststore "cs_chat_app_server/modules/request/store"
)

type getGroupRequestBiz struct {
	groupRepo    grouprepo.Repository
	notification notirepo.NotificationServiceRepository
}

func NewGetGroupRequestBiz(groupRepo grouprepo.Repository) *getGroupRequestBiz {
	return &getGroupRequestBiz{groupRepo: groupRepo}
}

// GetRequest send a group invitation request to user.
func (biz *getGroupRequestBiz) GetRequest(ctx context.Context, requesterId string, filter groupmdl.Filter) ([]requestmdl.Request, error) {

	var groupFilterFilter map[string]interface{}
	if filter == groupmdl.Sent {
		groupFilterFilter = requeststore.GetRequestSenderIdFilter(requesterId)
	} else {
		groupFilterFilter = requeststore.GetRequestReceiverIdFilter(requesterId)
	}

	ft := common.GetAndFilter(
		groupFilterFilter,
		requeststore.GetTypeFilterFilter(true),
	)

	requests, err := biz.groupRepo.FindRequests(ctx, ft)

	if err != nil {
		return nil, err
	}

	return requests, nil
}
