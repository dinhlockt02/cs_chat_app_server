package groupstore

import (
	"cs_chat_app_server/common"
	groupmdl "cs_chat_app_server/modules/group/model"
)

func GetGroupIdInIdListFilter(ids ...string) map[string]interface{} {
	mongoIds := make([]interface{}, 0, len(ids))
	for i := range ids {
		mongoId, err := common.ToObjectId(ids[i])
		if err == nil {
			mongoIds = append(mongoIds, mongoId)
		}
	}
	return common.GetInFilter("_id", mongoIds...)
}

func GetMemberIdInGroupMembersFilter(id string) map[string]interface{} {
	return map[string]interface{}{
		"members.id": id,
	}

}

func GetUserIdInIdListFilter(ids ...string) map[string]interface{} {
	mongoIds := make([]interface{}, 0, len(ids))
	for i := range ids {
		mongoId, err := common.ToObjectId(ids[i])
		if err == nil {
			mongoIds = append(mongoIds, mongoId)
		}
	}
	return common.GetInFilter("_id", mongoIds...)
}

func GetTypeFilter(typ groupmdl.GroupType) map[string]interface{} {
	return map[string]interface{}{
		"type": typ,
	}
}
