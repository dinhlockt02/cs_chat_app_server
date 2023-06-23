package groupbiz

import (
	"context"
	"cs_chat_app_server/common"
	groupmdl "cs_chat_app_server/modules/group/model"
	grouprepo "cs_chat_app_server/modules/group/repository"
	"errors"
)

type createGroupBiz struct {
	groupRepo grouprepo.Repository
}

func NewCreateGroupBiz(groupRepo grouprepo.Repository) *createGroupBiz {
	return &createGroupBiz{groupRepo: groupRepo}
}

// Create creates a group and add requester as a member.
func (biz *createGroupBiz) Create(ctx context.Context, requester string, data *groupmdl.Group) error {

	if err := data.Process(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	data.Type = groupmdl.TypeGroup
	data.Active = common.GetPointer(true)

	userFilter := make(map[string]interface{})
	_ = common.AddIdFilter(userFilter, requester)
	user, err := biz.groupRepo.FindUser(ctx, userFilter)
	if err != nil {
		return err
	}
	if user == nil {
		return common.ErrEntityNotFound("User", errors.New("user not found"))
	}

	data.Members = []groupmdl.GroupUser{groupmdl.GroupUser{
		Id:     requester,
		Name:   user.Name,
		Avatar: user.Avatar,
	}}

	if err = biz.groupRepo.CreateGroup(ctx, data); err != nil {
		return err
	}

	user.Groups = append(user.Groups, *data.Id)

	if err = biz.groupRepo.UpdateUser(ctx, userFilter, user); err != nil {
		return err
	}

	return nil
}
