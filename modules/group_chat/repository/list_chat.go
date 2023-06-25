package gchatrepo

import (
	"context"
	"cs_chat_app_server/common"
	gchatmdl "cs_chat_app_server/modules/group_chat/model"
	"errors"
)

// TODO: Cache User Data

func (repo *GroupChatRepository) List(
	ctx context.Context,
	filter map[string]interface{},
	paging gchatmdl.Paging,
) ([]gchatmdl.GroupChatItem, error) {
	list, err := repo.groupChatStore.List(ctx, filter, &paging)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	if len(list) == 0 {
		return nil, nil
	}

	idFilter, err := common.GetIdFilter(list[0].GroupId)
	group, err := repo.groupChatStore.FindGroup(ctx, idFilter)

	var users = make(map[string]*gchatmdl.User)

	for i, _ := range list {
		if _, ok := users[list[i].SenderId]; !ok {
			filter = map[string]interface{}{}
			err = common.AddIdFilter(filter, list[i].SenderId)
			if err != nil {
				return nil, err
			}
			user, err := repo.groupChatStore.FindUser(ctx, filter)
			if err != nil {
				return nil, err
			}
			if user == nil {
				return nil, common.ErrInternal(errors.New("User not found"))
			}
			users[list[i].SenderId] = user
		}

		list[i].Sender = users[list[i].SenderId]
		list[i].Group = group
	}
	return list, nil
}
