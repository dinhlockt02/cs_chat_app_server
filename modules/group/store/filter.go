package groupstore

import "cs_chat_app_server/common"

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
		"members": id,
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
