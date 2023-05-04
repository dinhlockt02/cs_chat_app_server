package pchatstore

import (
	"context"
	"cs_chat_app_server/common"
	pchatmdl "cs_chat_app_server/modules/personal_chat/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *mongoStore) List(
	ctx context.Context,
	filter map[string]interface{},
	paging common.Paging[string]) ([]pchatmdl.PersonalChatItem, error) {

	var value interface{} = paging.Value
	var err error

	if paging.Field == "id" {
		paging.Field = "_id"
		value, err = common.ToObjectId(paging.Value)
		if err != nil {
			return nil, common.ErrInvalidRequest(err)
		}
	}

	sort := 1
	filter[paging.Field] = map[string]interface{}{
		"$gt": value,
	}
	if paging.Order == "desc" {
		sort = -1
		filter[paging.Field] = map[string]interface{}{
			"$lt": paging.Value,
		}
	}

	opts := options.Find().SetLimit(paging.Limit).SetSort(bson.D{{paging.Field, sort}})

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
