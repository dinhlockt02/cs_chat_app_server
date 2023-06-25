package groupstore

import (
	"context"
	"cs_chat_app_server/common"
	groupmdl "cs_chat_app_server/modules/group/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *mongoStore) List(
	ctx context.Context,
	filter map[string]interface{},
) ([]groupmdl.Group, error) {

	// Hardcoded sort order
	opts := options.Find().SetSort(bson.D{{"latest_message.created_at", -1}})
	cursor, err := s.database.Collection(groupmdl.Group{}.CollectionName()).
		Find(ctx, filter, opts)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	var groups []groupmdl.Group
	if err = cursor.All(ctx, &groups); err != nil {
		return nil, common.ErrInternal(err)
	}
	return groups, nil
}
