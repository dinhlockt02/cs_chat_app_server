package subscriber

import (
	"context"
	"cs_chat_app_server/components/appcontext"
)

func Setup(appCtx appcontext.AppContext, ctx context.Context) {
	UpdateRequestWhenUserUpdateProfile(appCtx, ctx)
	UpdateRequestWhenGroupUpdated(appCtx, ctx)
	NotifyUserWhenNewGroupMessageReceived(appCtx, ctx)
	UpdateNotificationWhenUserUpdateProfile(appCtx, ctx)
	UpdateGroupsWhenUserUpdateProfile(appCtx, ctx)
	UpdateGroupLatestMessageWhenNewGroupMessageReceived(appCtx, ctx)
	UpdateNotificationWhenGroupUpdated(appCtx, ctx)
	DeleteNotificationWhenUserAcceptFriendRequest(appCtx, ctx)
	DeleteNotificationWhenUserDenyFriendRequest(appCtx, ctx)
	DeleteNotificationWhenUserAcceptGroupRequest(appCtx, ctx)
	DeleteNotificationWhenUserDenyGroupRequest(appCtx, ctx)
}
