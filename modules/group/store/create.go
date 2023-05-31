package groupstore

import (
	"context"
	"cs_chat_app_server/common"
	groupmdl "cs_chat_app_server/modules/group/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *mongoStore) Create(ctx context.Context, group *groupmdl.Group) error {
	result, err := s.database.Collection(group.CollectionName()).InsertOne(ctx, group)
	if err != nil {
		return common.ErrInternal(err)
	}
	createdId := result.InsertedID.(primitive.ObjectID).Hex()
	group.Id = &createdId
	return nil
}
