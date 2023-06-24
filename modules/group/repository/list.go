package grouprepo

import (
	"context"
	"cs_chat_app_server/common"
	groupmdl "cs_chat_app_server/modules/group/model"
	groupstore "cs_chat_app_server/modules/group/store"
)

func (repo *groupRepository) List(ctx context.Context, requester string, groupFilter map[string]interface{}) ([]groupmdl.Group, error) {
	groups, err := repo.groupStore.List(ctx, common.GetAndFilter(
		groupFilter,
		groupstore.GetActiveFilter(true),
	),
	)
	if err != nil {
		return nil, err
	}

	for i := range groups {
		if groups[i].Type != groupmdl.TypePersonal {
			continue
		}

		var friend *groupmdl.GroupUser
		if groups[i].Members[0].Id == requester {
			friend = &groups[i].Members[1]
		} else {
			friend = &groups[i].Members[0]
		}

		groups[i].Name = friend.Name
		groups[i].ImageUrl = &friend.Avatar
	}

	return groups, nil
}
