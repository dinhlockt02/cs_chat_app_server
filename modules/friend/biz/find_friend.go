package friendbiz

import (
	"context"
	"cs_chat_app_server/common"
	friendmodel "cs_chat_app_server/modules/friend/model"
	groupmdl "cs_chat_app_server/modules/group/model"
	grouprepo "cs_chat_app_server/modules/group/repository"
	groupstore "cs_chat_app_server/modules/group/store"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (biz *findFriendBiz) FindFriend(ctx context.Context, userId string) ([]friendmodel.FriendUser, error) {
	id, _ := primitive.ObjectIDFromHex(userId)
	user, err := biz.friendStore.FindUser(ctx, map[string]interface{}{
		"_id": id,
	})

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, common.ErrEntityNotFound("User", errors.New("user not found"))
	}

	var ids = make([]primitive.ObjectID, 0, len(user.Friends))
	for _, friend := range user.Friends {
		id, err = primitive.ObjectIDFromHex(friend)
		if err != nil {
			return nil, common.ErrInternal(err)
		}
		ids = append(ids, id)
	}
	friends, err := biz.friendStore.FindFriend(ctx, map[string]interface{}{
		"_id": map[string]interface{}{
			"$in": ids,
		},
	})

	for i := range friends {
		filter := common.GetAndFilter(
			groupstore.GetMemberIdInGroupMembersFilter(*friends[i].Id),
			groupstore.GetMemberIdInGroupMembersFilter(userId),
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
