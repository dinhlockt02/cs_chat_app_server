package notimodel

import (
	"cs_chat_app_server/common"
	"encoding/json"
	"fmt"
	"time"
)

type NotificationId byte

const (
	AcceptFriendRequestId NotificationId = iota + 1
	ReceiveFriendRequestId
)

type ChannelKey string

const (
	BasicChannel ChannelKey = "basic_channel"
)

type ActionKey string

const (
	Accept ActionKey = "accept"
	Deny             = "deny"
)

type NotificationObjectType string

const (
	User NotificationObjectType = "user"
)

type NotificationActionType string

const (
	AcceptRequest        NotificationActionType = "accept-request"
	ReceiveFriendRequest                        = "send-friend-request"
)

type Notification struct {
	common.MongoId        `bson:",inline" json:",inline"`
	common.MongoCreatedAt `json:",inline" bson:",inline"`

	// Owner is a string that represent the id of the user who receive notification
	Owner string `bson:"owner,omitempty" json:"owner,omitempty"`

	// Subject is a NotificationObject that represent the object do the Action
	Subject *NotificationObject `bson:"subject,omitempty" json:"subject,omitempty"`

	// Direct is a NotificationObject that represent the object was directly affected by the Action
	Direct *NotificationObject `bson:"direct,omitempty" json:"direct,omitempty"`

	// Indirect is a NotificationObject that represent the object was indirectly affected by the Action
	Indirect *NotificationObject `bson:"indirect,omitempty" json:"indirect,omitempty"`

	// Indirect is a NotificationObject that represent the object was appear in the action with a prep (in, for, of)
	Prep *NotificationObject `bson:"prep,omitempty" json:"prep,omitempty"`

	// Action is a string has type of NotificationActionType
	Action NotificationActionType `json:"action,omitempty" bson:"action,omitempty"`
}

func (Notification) CollectionName() string {
	return "notifications"
}

type notificationBuilder struct {
	Owner    string
	Subject  *NotificationObject
	Direct   *NotificationObject
	Indirect *NotificationObject
	Prep     *NotificationObject
	Action   NotificationActionType
}

func NewNotificationBuilder(action NotificationActionType, owner string) *notificationBuilder {
	result := new(notificationBuilder)
	result.Action = action
	result.Owner = owner
	return result
}

func (builder *notificationBuilder) SetSubject(object *NotificationObject) *notificationBuilder {
	builder.Subject = object
	return builder
}

func (builder *notificationBuilder) SetDirect(object *NotificationObject) *notificationBuilder {
	builder.Direct = object
	return builder
}

func (builder *notificationBuilder) SetIndirect(object *NotificationObject) *notificationBuilder {
	builder.Indirect = object
	return builder
}

func (builder *notificationBuilder) SetPrep(object *NotificationObject) *notificationBuilder {
	builder.Prep = object
	return builder
}

func (builder *notificationBuilder) Build() *Notification {
	now := time.Now()
	return &Notification{
		MongoCreatedAt: common.MongoCreatedAt{CreatedAt: &now},
		Subject:        builder.Subject,
		Direct:         builder.Direct,
		Indirect:       builder.Indirect,
		Prep:           builder.Prep,
		Owner:          builder.Owner,
		Action:         builder.Action,
	}
}

// GetMessage is a function that will return 2 values respectively
// is title and the content of the notification.
func (n *Notification) GetMessage() (title string, body string) {
	switch n.Action {
	case AcceptRequest:
		return "Accept friend request", fmt.Sprintf("%s accept your friend request", n.Subject.Name)
	case ReceiveFriendRequest:
		return "Friend request received", fmt.Sprintf("%s want to be friend with you", n.Prep.Name)
	default:
		return "", ""
	}
}

// GetContent is a function that will a map that will meet the awesome_notification requirement.
func (n *Notification) GetContent() (map[string]interface{}, error) {

	title, body := n.GetMessage()

	marshaledNotification, err := json.Marshal(n)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	switch n.Action {
	case AcceptRequest:
		return map[string]interface{}{
			"id":                  AcceptFriendRequestId,
			"channelKey":          BasicChannel,
			"displayOnForeground": true,
			"displayOnBackground": true,
			"notificationLayout":  "Default",
			"showWhen":            true,
			"autoDismissible":     true,
			"largeIcon":           n.Subject.Image,
			"privacy":             "Private",
			"payload": map[string]string{
				"notification": string(marshaledNotification),
			},
			"category": "Social",
			"title":    title,
			"body":     body,
			"locked":   false,
		}, nil
	case ReceiveFriendRequest:
		return map[string]interface{}{
			"id":                  ReceiveFriendRequestId,
			"channelKey":          BasicChannel,
			"displayOnForeground": true,
			"displayOnBackground": true,
			"notificationLayout":  "Default",
			"showWhen":            true,
			"autoDismissible":     true,
			"largeIcon":           n.Prep.Image,
			"privacy":             "Private",
			"payload": map[string]string{
				"notification": string(marshaledNotification),
			},
			"category": "Social",
			"title":    title,
			"body":     body,
			"locked":   false,
		}, nil
	default:
		return nil, nil
	}
}

// GetActionButton is a method that will a slice of map
// which each item ís an action button.
func (n *Notification) GetActionButton() []map[string]interface{} {
	switch n.Action {
	case AcceptRequest:
		return nil
	case ReceiveFriendRequest:
		return []map[string]interface{}{
			{
				"key":             Accept,
				"label":           GetActionKeyLabel(Accept),
				"autoDismissible": true,
				"actionType":      "DismissAction",
			},
			{
				"key":               Deny,
				"label":             GetActionKeyLabel(Deny),
				"isDangerousOption": true,
				"autoDismissible":   true,
				"actionType":        "DismissAction",
			},
		}
	default:
		return nil
	}

}

func GetActionKeyLabel(key ActionKey) string {
	switch key {
	case Accept:
		return "Accept"
	case Deny:
		return "Deny"
	default:
		return ""
	}
}
