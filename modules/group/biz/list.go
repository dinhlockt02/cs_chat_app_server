package groupbiz

import (
	"context"
	"cs_chat_app_server/common"
	groupmdl "cs_chat_app_server/modules/group/model"
	grouprepo "cs_chat_app_server/modules/group/repository"
	groupstore "cs_chat_app_server/modules/group/store"
)

type listGroupBiz struct {
	groupRepo grouprepo.Repository
}

func NewListGroupBiz(groupRepo grouprepo.Repository) *createGroupBiz {
	return &createGroupBiz{groupRepo: groupRepo}
}

func (biz *createGroupBiz) List(ctx context.Context, requesterId string, filters ...map[string]interface{}) ([]groupmdl.Group, error) {

	//filter := make(map[string]interface{})
	//_ = common.AddIdFilter(filter, requesterId)
	//
	//user, err := biz.groupRepo.FindUser(ctx, filter)
	//if err != nil {
	//	return nil, err
	//}

	filter := groupstore.GetMemberIdInGroupMembersFilter(requesterId)

	for _, f := range filters {
		filter = common.GetAndFilter(filter, f)
	}

	return biz.groupRepo.List(ctx, requesterId, filter)

}
