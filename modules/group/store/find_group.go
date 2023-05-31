package groupstore

import (
	"context"
	"cs_chat_app_server/common"
	groupmdl "cs_chat_app_server/modules/group/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *mongoStore) FindGroup(
	ctx context.Context,
	filter map[string]interface{},
) (*groupmdl.Group, error) {

	var group *groupmdl.Group

	result := s.database.Collection(groupmdl.Group{}.CollectionName()).
		FindOne(ctx, filter)
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}
	}
	if err := result.Decode(&group); err != nil {
		return nil, common.ErrInternal(err)
	}
	return group, nil
}
