package groupbiz

import (
	"context"
	"cs_chat_app_server/common"
	notirepo "cs_chat_app_server/components/notification/repository"
	friendmodel "cs_chat_app_server/modules/friend/model"
	grouprepo "cs_chat_app_server/modules/group/repository"
	requeststore "cs_chat_app_server/modules/request/store"
)

type rejectGroupRequestBiz struct {
	groupRepo    grouprepo.Repository
	notification notirepo.NotificationRepository
}

func NewRejectGroupRequestBiz(groupRepo grouprepo.Repository) *rejectGroupRequestBiz {
	return &rejectGroupRequestBiz{groupRepo: groupRepo}
}

// RejectRequest send a group invitation request to user.
func (biz *rejectGroupRequestBiz) RejectRequest(ctx context.Context, requesterId string, groupId string) error {
	// Find exists request
	requesterFilter := requeststore.GetRequestReceiverIdFilter(requesterId)
	groupFilter := requeststore.GetRequestGroupIdFilter(groupId)
	ft := common.GetAndFilter(requesterFilter, groupFilter)
	existedRequest, err := biz.groupRepo.FindRequest(ctx, ft)
	if err != nil {
		return err
	}
	if existedRequest == nil {
		return common.ErrInvalidRequest(friendmodel.ErrRequestNotFound)
	}

	// Delete request
	filter := make(map[string]interface{})
	err = common.AddIdFilter(filter, *existedRequest.Id)
	err = biz.groupRepo.DeleteRequest(ctx, filter)
	if err != nil {
		return err
	}
	// TODO: send push notification new member joined
	return nil
}
