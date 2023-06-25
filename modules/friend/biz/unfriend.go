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

type UnfriendFriendStore interface {
	FindUser(ctx context.Context, filter map[string]interface{}) (*friendmodel.User, error)
	UpdateUser(ctx context.Context, filter map[string]interface{}, updatedUser *friendmodel.User) error
}

type unfriendBiz struct {
	friendStore     UnfriendFriendStore
	groupRepository grouprepo.Repository
}

func NewUnfriendBiz(
	friendStore UnfriendFriendStore,
	groupRepository grouprepo.Repository) *unfriendBiz {
	return &unfriendBiz{
		friendStore:     friendStore,
		groupRepository: groupRepository,
	}
}

func (biz *unfriendBiz) Unfriend(ctx context.Context, userId string, friendId string) error {
	id, _ := primitive.ObjectIDFromHex(userId)
	user, err := biz.friendStore.FindUser(ctx, map[string]interface{}{
		"_id": id,
	})

	if err != nil {
		return err
	}
	if user == nil {
		return common.ErrEntityNotFound("User", errors.New("user not found"))
	}
	for i := range user.Friends {
		if user.Friends[i] == friendId {
			user.Friends = append(user.Friends[:i], user.Friends[i+1:]...)
			err = biz.friendStore.UpdateUser(ctx, map[string]interface{}{
				"_id": id,
			}, user)
			if err != nil {
				return err
			}
			break
		}
	}

	id, _ = primitive.ObjectIDFromHex(friendId)
	friend, err := biz.friendStore.FindUser(ctx, map[string]interface{}{
		"_id": id,
	})
	if err != nil {
		return err
	}
	if friend == nil {
		return common.ErrEntityNotFound("User", errors.New("receiver not found"))
	}
	for i := range friend.Friends {
		if friend.Friends[i] == userId {
			friend.Friends = append(friend.Friends[:i], friend.Friends[i+1:]...)
			err = biz.friendStore.UpdateUser(ctx, map[string]interface{}{
				"_id": id,
			}, friend)
			if err != nil {
				return err
			}
			break
		}
	}

	// Set group inactive
	filter := common.GetAndFilter(
		groupstore.GetMemberIdInGroupMembersFilter(userId),
		groupstore.GetMemberIdInGroupMembersFilter(friendId),
		groupstore.GetTypeFilter(groupmdl.TypePersonal),
	)
	err = biz.groupRepository.UpdateGroup(ctx, filter, &groupmdl.UpdateGroup{
		Active: common.GetPointer(false),
	})
	if err != nil {
		return err
	}
	return nil
}
