package gchatmdl

import (
	"cs_chat_app_server/common"
	"time"
)

type MessageType string

const (
	image MessageType = "image"
	text  MessageType = "text"
	video MessageType = "video"
)

type GroupChatItem struct {
	common.MongoId        `bson:",inline" json:",inline,omitempty"`
	Type                  MessageType `json:"type" bson:"type"`
	SenderId              string      `bson:"sender" json:"-"`
	Sender                *User       `json:"sender" bson:"-"`
	GroupId               string      `json:"-" bson:"group"`
	Group                 *Group      `json:"group" bson:"-"`
	Message               string      `bson:"message" json:"message"`
	Optional              *string     `bson:"optional,omitempty" json:"optional,omitempty"`
	ThumbnailUrl          *string     `bson:"thumbnail_url,omitempty" json:"thumbnail_url,omitempty"`
	VideoUrl              *string     `bson:"video_url,omitempty" json:"video_url,omitempty"`
	ImageUrl              *string     `bson:"image_url,omitempty" json:"image_url,omitempty"`
	common.MongoCreatedAt `bson:",inline" json:",inline"`
	IsMe                  *bool `json:"is_me,omitempty" bson:"-"`
}

func (p *GroupChatItem) CollectionName() string {
	return common.GroupChatHistoryCollectionName
}

func (p *GroupChatItem) Process() error {
	now := time.Now()
	p.CreatedAt = &now
	p.Id = nil

	if p.Type != video {
		p.ThumbnailUrl = nil
		p.VideoUrl = nil
	}

	if p.Type != image {
		p.ImageUrl = nil
	}

	return nil
}
