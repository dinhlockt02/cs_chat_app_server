package friendstore

import (
	"context"
	"cs_chat_app_server/common"
	friendmodel "cs_chat_app_server/modules/friend/model"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *mongoStore) FindRequests(ctx context.Context, filter map[string]interface{}) ([]friendmodel.Request, error) {
	var request []friendmodel.Request
	log.Info().Msgf("%v", filter)
	cursor, err := s.database.Collection(friendmodel.Request{}.CollectionName()).Find(ctx, filter)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	if err = cursor.All(ctx, &request); err != nil {
		return nil, common.ErrInternal(err)
	}

	return request, nil
}

func (s *mongoStore) FindRequest(ctx context.Context, userId string, otherId string) (*friendmodel.Request, error) {
	var request friendmodel.Request
	query := bson.D{{
		"$or", []interface{}{
			bson.D{{"sender.id", userId}, {"receiver.id", otherId}},
			bson.D{{"sender.id", otherId}, {"receiver.id", userId}},
		},
	}}
	result := s.database.Collection(request.CollectionName()).FindOne(ctx, query)
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
