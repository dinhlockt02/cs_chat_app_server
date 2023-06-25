package friendstore

import (
	"context"
	"cs_chat_app_server/common"
	friendmodel "cs_chat_app_server/modules/friend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindFriend is a method for finding user by filter
// and store data into a slice of FriendUser struct
func (s *mongoStore) FindFriend(ctx context.Context, filter map[string]interface{}) ([]friendmodel.FriendUser, error) {
	var friends []friendmodel.FriendUser

	// Hardcoded sort order
	opts := options.Find().SetSort(bson.D{{"name", 1}})
	cur, err := s.database.Collection(friendmodel.FriendUser{}.CollectionName()).Find(ctx, filter, opts)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	if err = cur.All(ctx, &friends); err != nil {
		return nil, common.ErrInternal(err)
	}

	return friends, nil
}
