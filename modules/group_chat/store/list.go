package gchatstore

import (
	"context"
	"cs_chat_app_server/common"
	gchatmdl "cs_chat_app_server/modules/group_chat/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *mongoStore) List(
	ctx context.Context,
	filter map[string]interface{},
	paging gchatmdl.Paging,
) ([]gchatmdl.GroupChatItem, error) {

	paging.Process()

	if paging.LastId != nil {
		id, err := primitive.ObjectIDFromHex(*paging.LastId)
		if err != nil {
			return nil, common.ErrInvalidRequest(err)
		}

		if *paging.Order == gchatmdl.ASC {
			filter["_id"] = map[string]interface{}{
				"$gt": id,
			}
		} else {
			filter["_id"] = map[string]interface{}{
				"$lt": id,
			}
		}
	}
	var sort = -1
	if *paging.Order == gchatmdl.ASC {
		sort = 1
	}

	opts := options.Find().SetLimit(*paging.Limit).SetSort(bson.D{{"_id", sort}})

	cursor, err := s.database.
		Collection((&gchatmdl.GroupChatItem{}).CollectionName()).
		Find(ctx, filter, opts)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	var rs []gchatmdl.GroupChatItem
	err = cursor.All(ctx, &rs)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return rs, nil
}
