package pchatstore

import (
	"context"
	"cs_chat_app_server/common"
	pchatmdl "cs_chat_app_server/modules/personal_chat/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *mongoStore) List(
	ctx context.Context,
	filter map[string]interface{},
	paging pchatmdl.Paging,
) ([]pchatmdl.PersonalChatItem, error) {

	paging.Process()

	if paging.LastId != nil {
		id, err := primitive.ObjectIDFromHex(*paging.LastId)
		if err != nil {
			return nil, common.ErrInvalidRequest(err)
		}

		if *paging.Order == pchatmdl.ASC {
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
	if *paging.Order == pchatmdl.ASC {
		sort = 1
	}

	opts := options.Find().SetLimit(*paging.Limit).SetSort(bson.D{{"_id", sort}})

	cursor, err := s.database.
		Collection((&pchatmdl.PersonalChatItem{}).CollectionName()).
		Find(ctx, filter, opts)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	var rs []pchatmdl.PersonalChatItem
	err = cursor.All(ctx, &rs)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return rs, nil
}
