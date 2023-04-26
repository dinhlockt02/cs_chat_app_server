package authmiddleware

import (
	"context"
	"cs_chat_app_server/common"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoStore struct {
	database *mongo.Database
}

func NewMongoStore(database *mongo.Database) *mongoStore {
	return &mongoStore{database: database}
}

func (s *mongoStore) FindOne(ctx context.Context, filter map[string]interface{}) (*User, error) {
	var user User
	result := s.database.Collection(user.CollectionName()).FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, common.ErrInternal(err)
	}
	if err := result.Decode(&user); err != nil {
		return nil, common.ErrInternal(err)
	}
	return &user, nil
}
