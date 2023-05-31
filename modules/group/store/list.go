package groupstore

import (
	"context"
	"cs_chat_app_server/common"
	groupmdl "cs_chat_app_server/modules/group/model"
)

func (s *mongoStore) List(
	ctx context.Context,
	filter map[string]interface{},
) ([]groupmdl.Group, error) {

	cursor, err := s.database.Collection(groupmdl.Group{}.CollectionName()).
		Find(ctx, filter)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	var groups []groupmdl.Group
	if err = cursor.All(ctx, &groups); err != nil {
		return nil, common.ErrInternal(err)
	}
	return groups, nil
}
