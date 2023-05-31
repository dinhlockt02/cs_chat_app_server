package requeststore

import (
	"context"
	"cs_chat_app_server/common"
	requestmdl "cs_chat_app_server/modules/request/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *mongoStore) UpdateRequests(
	ctx context.Context,
	filter map[string]interface{},
	data *requestmdl.UpdateRequest,
) error {
	updateData := bson.D{{
		"$set", data,
	}}
	_, err := s.database.Collection(data.CollectionName()).UpdateMany(ctx, filter, updateData)
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
