package groupstore

import (
	"context"
	"cs_chat_app_server/common"
	groupmdl "cs_chat_app_server/modules/group/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *mongoStore) UpdateUser(ctx context.Context, filter map[string]interface{}, updatedUser *groupmdl.User) error {
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
