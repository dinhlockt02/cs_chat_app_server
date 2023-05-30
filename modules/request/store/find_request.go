package requeststore

import (
	"context"
	"cs_chat_app_server/common"
	requestmdl "cs_chat_app_server/modules/request/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *mongoStore) FindRequests(ctx context.Context, filter map[string]interface{}) ([]requestmdl.Request, error) {
	var request []requestmdl.Request
	cursor, err := s.database.Collection(requestmdl.Request{}.CollectionName()).Find(ctx, filter)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	if err = cursor.All(ctx, &request); err != nil {
		return nil, common.ErrInternal(err)
	}

	return request, nil
}

func (s *mongoStore) FindRequest(ctx context.Context, filter map[string]interface{}) (*requestmdl.Request, error) {
	var request requestmdl.Request
	result := s.database.Collection(request.CollectionName()).FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, common.ErrInternal(err)
	}

	if err := result.Decode(&request); err != nil {
		return nil, common.ErrInternal(err)
	}
	return &request, nil
}
