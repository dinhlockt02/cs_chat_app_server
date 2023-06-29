package groupbiz

import (
	"context"
	"cs_chat_app_server/common"
	notirepo "cs_chat_app_server/components/notification/repository"
	groupmdl "cs_chat_app_server/modules/group/model"
	grouprepo "cs_chat_app_server/modules/group/repository"
	"errors"
)

type leaveGroupBiz struct {
	groupRepo    grouprepo.Repository
	notification notirepo.NotificationServiceRepository
}

func NewLeaveGroupBiz(groupRepo grouprepo.Repository) *leaveGroupBiz {
	return &leaveGroupBiz{groupRepo: groupRepo}
}

// Leave present a requester leave group.
func (biz *leaveGroupBiz) Leave(ctx context.Context, requesterId string, groupId string) error {
	// Find sender
	filter, err := common.GetIdFilter(requesterId)
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
	var groupIndex = -1
	for i, gid := range requester.Groups {
		if gid == groupId {
			groupIndex = i
		}
	}
	if groupIndex != -1 {
		requester.Groups = append(requester.Groups[:groupIndex], requester.Groups[groupIndex+1:]...)
		filter = make(map[string]interface{})
		filter, err = common.GetIdFilter(requesterId)
		if err != nil {
			return err
		}

		err = biz.groupRepo.UpdateUser(ctx, filter, requester)
		if err != nil {
			return err
		}
	}

	// Update Group
	var memberIndex = -1
	for i, mid := range group.Members {
		if mid.Id == requesterId {
			memberIndex = i
		}
	}
	if memberIndex != -1 {
		updateGroup := &groupmdl.UpdateGroup{}
		updateGroup.Members = append(group.Members[:memberIndex], group.Members[memberIndex+1:]...)
		if len(updateGroup.Members) == 0 {
			updateGroup.Active = common.GetPointer(false)
		}
		filter = make(map[string]interface{})
		err = common.AddIdFilter(filter, groupId)
		if err != nil {
			return err
		}
		err = biz.groupRepo.UpdateGroup(ctx, filter, updateGroup)
		if err != nil {
			return err
		}
	}

	return nil
}
