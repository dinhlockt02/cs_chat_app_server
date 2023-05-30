package notirepo

import (
	"context"
	notimodel "cs_chat_app_server/components/notification/model"
)

type NotificationRepository interface {
	// CreateAcceptFriendNotification is a method that will create, store and push notification
	//
	// It should be used when the subject accept the indirect (aka owner)'s friend request
	CreateAcceptFriendNotification(
		ctx context.Context,
		owner string,
		subject *notimodel.NotificationObject,
		indirect *notimodel.NotificationObject,
	) error

	// CreateReceiveFriendRequestNotification is a method that will create, store and push notification
	//
	// It should be used when the Subject (aka owner) received the friend request from Prep's
	CreateReceiveFriendRequestNotification(
		ctx context.Context,
		owner string,
		subject *notimodel.NotificationObject,
		prep *notimodel.NotificationObject,
	) error
}
