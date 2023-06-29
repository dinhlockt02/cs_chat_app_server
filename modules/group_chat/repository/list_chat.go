package gchatrepo

import (
	"context"
	"cs_chat_app_server/common"
	gchatmdl "cs_chat_app_server/modules/group_chat/model"
	"errors"
	"github.com/rs/zerolog/log"
)

// TODO: Cache User Data

func (repo *GroupChatRepository) List(
	ctx context.Context,
	filter map[string]interface{},
	paging gchatmdl.Paging,
) ([]gchatmdl.GroupChatItem, error) {
	list, err := repo.groupChatStore.List(ctx, filter, paging)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	if len(list) == 0 {
		return nil, nil
	}

	groupMap := make(map[string]*gchatmdl.Group)

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

		if _, ok := groupMap[list[i].GroupId]; !ok {
			idFilter, err := common.GetIdFilter(list[i].GroupId)
			if err != nil {
				log.Error().Err(err).Str("package", "gchatrepo.List").Send()
				return nil, common.ErrInternal(err)
			}
			group, err := repo.groupChatStore.FindGroup(ctx, idFilter)
			if err != nil {
				log.Error().Err(err).Str("package", "gchatrepo.List").Send()
				return nil, common.ErrInternal(err)
			}
			if group == nil {
				log.Error().Err(err).Str("package", "gchatrepo.List").Send()
				return nil, common.ErrInternal(errors.New("group not found"))
			}
			groupMap[list[i].GroupId] = group
		}
		list[i].Group = groupMap[list[i].GroupId]
	}
	return list, nil
}
