package groupbiz

import (
	"context"
	"cs_chat_app_server/common"
	notirepo "cs_chat_app_server/components/notification/repository"
	friendmodel "cs_chat_app_server/modules/friend/model"
	grouprepo "cs_chat_app_server/modules/group/repository"
	requeststore "cs_chat_app_server/modules/request/store"
	"errors"
)

type acceptGroupRequestBiz struct {
	groupRepo    grouprepo.Repository
	notification notirepo.NotificationRepository
}

func NewAcceptGroupRequestBiz(groupRepo grouprepo.Repository) *acceptGroupRequestBiz {
	return &acceptGroupRequestBiz{groupRepo: groupRepo}
}

// AcceptRequest send a group invitation request to user.
func (biz *acceptGroupRequestBiz) AcceptRequest(ctx context.Context, requesterId string, groupId string) error {
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

	// Find sender
	filter := make(map[string]interface{})
	err = common.AddIdFilter(filter, requesterId)
	requester, err := biz.groupRepo.FindUser(ctx, filter)
	if err != nil {
		return err
	}
	if requester == nil {
		return common.ErrEntityNotFound("User", errors.New("sender not found"))
	}

	// Find Group
	filter = make(map[string]interface{})
	err = common.AddIdFilter(filter, groupId)
	group, err := biz.groupRepo.FindGroup(ctx, filter)
	if group == nil {
		return common.ErrEntityNotFound("Group", errors.New("group not found"))
	}

	// Update Requester
	requester.Groups = append(requester.Groups, groupId)
	filter = make(map[string]interface{})
	err = common.AddIdFilter(filter, requesterId)
	if err != nil {
		return err
	}

	err = biz.groupRepo.UpdateUser(ctx, filter, requester)
	if err != nil {
		return err
	}

	// Update Group
	group.Members = append(group.Members, requesterId)
	group.Active = common.GetPointer(true)
	filter = make(map[string]interface{})
	err = common.AddIdFilter(filter, groupId)
	if err != nil {
		return err
	}
	err = biz.groupRepo.UpdateGroup(ctx, filter, group)
	if err != nil {
		return err
	}

	// Delete request
	filter = make(map[string]interface{})
	err = common.AddIdFilter(filter, *existedRequest.Id)
	err = biz.groupRepo.DeleteRequest(ctx, filter)
	if err != nil {
		return err
	}
	// TODO: send push notification new member joined
	return nil
}
