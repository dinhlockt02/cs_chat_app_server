package gchatmdl

import (
	"cs_chat_app_server/common"
	"time"
)

type MessageType string

const (
	Image  MessageType = "image"
	Video              = "video"
	Record             = "record"
	File               = "file"
	System             = "system"
)

type SystemEvent string

const (
	GroupCreated SystemEvent = "group-created"
	MemberJoined SystemEvent = "member-joined"
	MemberLeaved SystemEvent = "member-leaved"
)

type GroupChatItem struct {
	common.MongoId        `bson:",inline" json:",inline,omitempty"`
	Type                  MessageType  `json:"type" bson:"type"`
	SenderId              string       `bson:"sender" json:"-"`
	Sender                *User        `json:"sender" bson:"-"`
	GroupId               string       `json:"-" bson:"group"`
	Group                 *Group       `json:"group" bson:"-"`
	Message               string       `bson:"message" json:"message"`
	Optional              *string      `bson:"optional,omitempty" json:"optional,omitempty"`
	ThumbnailUrl          *string      `bson:"thumbnail_url,omitempty" json:"thumbnail_url,omitempty"`
	VideoUrl              *string      `bson:"video_url,omitempty" json:"video_url,omitempty"`
	ImageUrl              *string      `bson:"image_url,omitempty" json:"image_url,omitempty"`
	RecordUrl             *string      `bson:"record_url,omitempty" json:"record_url,omitempty"`
	SizeImage             *int64       `bson:"size_image,omitempty" json:"size_image,omitempty"`
	SystemEvent           *SystemEvent `json:"system_event,omitempty" bson:"system_event,omitempty"`
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

	if p.Type != Video {
		p.ThumbnailUrl = nil
		p.VideoUrl = nil
	}

	if p.Type != Image {
		p.ImageUrl = nil
	}

	if p.Type != Record {
		p.RecordUrl = nil
	}

	if p.Type != File {
		p.SizeImage = nil
	}

	return nil
}
