package friendstore

import (
	"context"
	"cs_chat_app_server/common"
	friendmodel "cs_chat_app_server/modules/friend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *mongoStore) DeleteRequest(ctx context.Context, requestId string) error {
	id, _ := primitive.ObjectIDFromHex(requestId)
	_, err := s.database.Collection(friendmodel.Request{}.CollectionName()).DeleteOne(ctx, bson.D{
		{"_id", id},
	})
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
