package friendstore

import (
	"context"
	"cs_chat_app_server/common"
	friendmodel "cs_chat_app_server/modules/friend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *mongoStore) UpdateUser(ctx context.Context, filter map[string]interface{}, updatedUser *friendmodel.User) error {
	updatedUser.Id = nil
	updateData := bson.D{{
		"$set", updatedUser,
	}}
	_, err := s.database.
		Collection(updatedUser.CollectionName()).
		UpdateOne(ctx, filter, updateData)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		return common.ErrInternal(err)
	}
	return nil
}

func (s *mongoStore) NewFilterById(id string) (map[string]interface{}, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, common.ErrInvalidRequest(err)
	}
	return map[string]interface{}{
		"_id": _id,
	}, nil
}
