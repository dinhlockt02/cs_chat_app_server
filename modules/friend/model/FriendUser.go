package friendmodel

import "time"

type FriendUser struct {
	Id         *string    `json:"id" bson:"_id,omitempty"`
	Avatar     string     `json:"avatar" bson:"avatar"`
	Name       string     `json:"name" bson:"name"`
	LastSeenAt *time.Time `json:"last_seen_at" bson:"last_seen_at"`
}

func (FriendUser) CollectionName() string {
	return "users"
}
