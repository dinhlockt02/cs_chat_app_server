package groupstore

import (
	"context"
	"cs_chat_app_server/common"
	groupmdl "cs_chat_app_server/modules/group/model"
)

func (s *mongoStore) DeleteGroups(ctx context.Context, filter map[string]interface{}) error {
	_, err := s.database.Collection(groupmdl.Group{}.CollectionName()).DeleteMany(ctx, filter)
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
