package groupbiz

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/pubsub"
	groupmdl "cs_chat_app_server/modules/group/model"
	grouprepo "cs_chat_app_server/modules/group/repository"
	groupstore "cs_chat_app_server/modules/group/store"
)

type updateGroupBiz struct {
	groupRepo grouprepo.Repository
	ps        pubsub.PubSub
}

func NewUpdateGroupBiz(
	groupRepo grouprepo.Repository,
	ps pubsub.PubSub,
) *updateGroupBiz {
	return &updateGroupBiz{
		groupRepo: groupRepo,
		ps:        ps,
	}
}

// Update allows requester, who is a member of group, update group info.
// If requester is not a group member, no error will be thrown
//
// Side effect: Update publish a common.TopicGroupUpdated
func (biz *updateGroupBiz) Update(ctx context.Context, requester string, groupId string, group *groupmdl.UpdateGroup) error {

	filter := make(map[string]interface{})
	err := common.AddIdFilter(filter, groupId)

	if err != nil {
		return common.ErrInvalidRequest(err)
	}

	filter = common.GetAndFilter(filter, groupstore.GetMemberIdInGroupMembersFilter(requester))

	err = biz.groupRepo.UpdateGroup(ctx, filter, group)
	if err != nil {
		return err
	}
	_ = biz.ps.Publish(ctx, common.TopicGroupUpdated, groupId)

	return nil
}
