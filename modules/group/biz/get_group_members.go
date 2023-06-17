package groupbiz

import (
	"context"
	groupmdl "cs_chat_app_server/modules/group/model"
	grouprepo "cs_chat_app_server/modules/group/repository"
	groupstore "cs_chat_app_server/modules/group/store"
)

type getGroupMembersBiz struct {
	groupRepo grouprepo.Repository
}

func NewGetGroupMembersBiz(groupRepo grouprepo.Repository) *getGroupBiz {
	return &getGroupBiz{groupRepo: groupRepo}
}

func (biz *getGroupBiz) GetGroupUsers(ctx context.Context, userIds ...string) ([]groupmdl.User, error) {
	return biz.groupRepo.FindUsers(ctx, groupstore.GetUserIdInIdListFilter(userIds...))
}
