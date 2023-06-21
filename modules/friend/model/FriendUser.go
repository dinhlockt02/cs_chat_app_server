package friendmodel

type FriendUser struct {
	Id     *string `json:"id" bson:"_id,omitempty"`
	Avatar string  `json:"avatar" bson:"avatar"`
	Name   string  `json:"name" bson:"name"`
	Group  string  `json:"group"`
}

func (FriendUser) CollectionName() string {
	return "users"
}
