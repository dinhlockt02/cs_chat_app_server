package groupstore

import (
	"context"
	"cs_chat_app_server/common"
	groupmdl "cs_chat_app_server/modules/group/model"
)

func (s *mongoStore) FindUsers(ctx context.Context, filter map[string]interface{}) ([]groupmdl.User, error) {
	var users []groupmdl.User
	cursor, err := s.database.Collection(groupmdl.User{}.CollectionName()).Find(ctx, filter)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	if err = cursor.All(ctx, &users); err != nil {
		return nil, common.ErrInternal(err)
	}

	return users, nil
}
