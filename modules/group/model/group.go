package groupmdl

import (
	"cs_chat_app_server/common"
	"errors"
)

type GroupType string

type GroupUser struct {
	Id     string `json:"id" bson:"id"`
	Name   string `json:"name" bson:"name"`
	Avatar string `bson:"avatar" json:"avatar"`
}

type GroupMessage struct {
	Message               string `bson:"message" json:"message"`
	SenderId              string `json:"sender_id" bson:"sender_id"`
	SenderName            string `json:"sender_name" bson:"sender_name"`
	common.MongoCreatedAt `bson:",inline" json:",inline"`
}

const (
	TypePersonal GroupType = "personal"
	TypeGroup              = "group"
)

type Group struct {
	common.MongoId `json:",inline,omitempty" bson:",inline,omitempty"`
	Name           string        `bson:"name" json:"name"`
	Members        []GroupUser   `bson:"members,omitempty" json:"members,omitempty"`
	InvitedUsers   []string      `json:"invited_users,omitempty" bson:"-"`
	ImageUrl       *string       `json:"image_url" bson:"image_url"`
	Type           GroupType     `bson:"type" json:"type"`
	Active         *bool         `json:"active,omitempty" bson:"active,omitempty"`
	LatestMessage  *GroupMessage `json:"latest_message,omitempty" bson:"latest_message,omitempty"`
}

func (Group) CollectionName() string {
	return common.GroupCollectionName
}

func (g *Group) Process() error {
	errs := common.ValidationError{}
	if g.ImageUrl != nil && !common.URLRegexp.Match([]byte(*g.ImageUrl)) {
		errs = append(errs, errors.New("invalid group image url"))
	}
	if len(g.Name) <= 0 {
		errs = append(errs, errors.New("invalid group name"))
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

type ListGroupQuery struct {
	Type *GroupType `form:"type"`
}

func (q *ListGroupQuery) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	if q.Type != nil {
		m["type"] = q.Type
	}

	return m
}
