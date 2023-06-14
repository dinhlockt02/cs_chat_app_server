package gchatstore

import (
	"context"
	"cs_chat_app_server/common"
	gchatmdl "cs_chat_app_server/modules/group_chat/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *mongoStore) FindGroup(ctx context.Context, filter map[string]interface{}) (*gchatmdl.Group, error) {
	result := s.database.
		Collection(gchatmdl.Group{}.CollectionName()).
		FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, common.ErrInternal(err)
	}
	var group = new(gchatmdl.Group)
	if err := result.Decode(&group); err != nil {
		return nil, err
	}
	return group, nil
}
