package groupbiz

import (
	"context"
	"cs_chat_app_server/common"
	notirepo "cs_chat_app_server/components/notification/repository"
	groupmdl "cs_chat_app_server/modules/group/model"
	grouprepo "cs_chat_app_server/modules/group/repository"
	"errors"
)

type getGroupBiz struct {
	groupRepo    grouprepo.Repository
	notification notirepo.NotificationServiceRepository
}

func NewGetGroupBiz(groupRepo grouprepo.Repository) *getGroupBiz {
	return &getGroupBiz{groupRepo: groupRepo}
}

// GetById returns a group by id.
func (biz *getGroupBiz) GetById(ctx context.Context, groupId string) (*groupmdl.Group, error) {

	// Find Group
	filter := make(map[string]interface{})
	err := common.AddIdFilter(filter, groupId)
	if err != nil {
		return nil, common.ErrInvalidRequest(errors.New("invalid group id"))
	}
	group, err := biz.groupRepo.FindGroup(ctx, filter)
	if group == nil {
		return nil, common.ErrEntityNotFound("Group", errors.New("group not found"))
	}

	return group, nil
}
