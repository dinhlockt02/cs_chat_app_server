package groupbiz

import (
	"context"
	"cs_chat_app_server/common"
	notirepo "cs_chat_app_server/components/notification/repository"
	friendmodel "cs_chat_app_server/modules/friend/model"
	grouprepo "cs_chat_app_server/modules/group/repository"
	requeststore "cs_chat_app_server/modules/request/store"
)

type recallGroupRequestBiz struct {
	groupRepo    grouprepo.Repository
	notification notirepo.NotificationServiceRepository
}

func NewRecallGroupRequestBiz(groupRepo grouprepo.Repository) *recallGroupRequestBiz {
	return &recallGroupRequestBiz{groupRepo: groupRepo}
}

// RecallRequest send a group invitation request to user.
func (biz *recallGroupRequestBiz) RecallRequest(ctx context.Context, requester string, user string, groupId string) error {
	// Find exists request
	senderFilter := requeststore.GetRequestSenderIdFilter(requester)
	receiverFilter := requeststore.GetRequestReceiverIdFilter(user)
	groupFilter := requeststore.GetRequestGroupIdFilter(groupId)
	ft := common.GetAndFilter(senderFilter, receiverFilter, groupFilter)
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
