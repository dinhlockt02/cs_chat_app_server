package friendstore

import (
	"context"
	"cs_chat_app_server/common"
	friendmodel "cs_chat_app_server/modules/friend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *mongoStore) CreateRequest(ctx context.Context, request *friendmodel.Request) error {
	result, err := s.database.Collection(request.CollectionName()).InsertOne(ctx, request)
	if err != nil {
		return common.ErrInternal(err)
	}
	insertedId := result.InsertedID.(primitive.ObjectID).Hex()
	request.Id = &insertedId
	return nil
}
