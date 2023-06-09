package friendbiz

import (
	"context"
	"cs_chat_app_server/common"
	friendmodel "cs_chat_app_server/modules/friend/model"
	groupmdl "cs_chat_app_server/modules/group/model"
	grouprepo "cs_chat_app_server/modules/group/repository"
	groupstore "cs_chat_app_server/modules/group/store"
	"errors"
)

type FindFriendFriendStore interface {
	FindFriend(ctx context.Context, filter map[string]interface{}) ([]friendmodel.FriendUser, error)
	FindUser(ctx context.Context, filter map[string]interface{}) (*friendmodel.User, error)
}

type findFriendBiz struct {
	friendStore     FindFriendFriendStore
	groupRepository grouprepo.Repository
}

func NewFindFriendBiz(
	friendStore FindFriendFriendStore,
	groupRepository grouprepo.Repository,
) *findFriendBiz {
	return &findFriendBiz{
		friendStore:     friendStore,
		groupRepository: groupRepository,
	}
}

func (biz *findFriendBiz) FindFriend(ctx context.Context, requesterId string, ft map[string]interface{}) ([]friendmodel.FriendUser, error) {
	filter, err := common.GetIdFilter(requesterId)
	if err != nil {
		return nil, err
	}
	user, err := biz.friendStore.FindUser(ctx, filter)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, common.ErrEntityNotFound("User", errors.New("user not found"))
	}

	filter = common.GetIdInIdListFilter(user.Friends...)
	friends, err := biz.friendStore.FindFriend(ctx, common.GetAndFilter(
		filter,
		ft,
	))

	for i := range friends {
		filter = common.GetAndFilter(
			groupstore.GetMemberIdInGroupMembersFilter(*friends[i].Id),
			groupstore.GetMemberIdInGroupMembersFilter(requesterId),
			groupstore.GetTypeFilter(groupmdl.TypePersonal),
		)
		group, err := biz.groupRepository.FindGroup(ctx, filter)
		if err != nil {
			return nil, err
		}

		friends[i].Group = *group.Id
	}
	if err != nil {
		return nil, err
	}
	return friends, nil
}
