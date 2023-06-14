package subscriber

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	groupstore "cs_chat_app_server/modules/group/store"
	gchatmdl "cs_chat_app_server/modules/group_chat/model"
	gchatstore "cs_chat_app_server/modules/group_chat/store"

	"github.com/rs/zerolog/log"
)

func NotifyUserWhenNewGroupMessageReceived(appCtx appcontext.AppContext, ctx context.Context) {
	ch := appCtx.PubSub().Subscribe(ctx, common.TopicNewGroupMessageCreated)

	groupStore := groupstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
	groupChatStore := gchatstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
	go func() {
		for messageId := range ch {
			go func(ctx context.Context, messageId string) {
				defer common.Recovery()

				// Extract message from messageId
				filter, err := common.GetIdFilter(messageId)
				if err != nil {
					panic(err)
				}

				message, err := groupChatStore.FindMessage(ctx, filter)

				if err != nil {
					panic(err)
				}

				// Extract group from groupId
				filter, err = common.GetIdFilter(message.GroupId)
				if err != nil {
					panic(err)
				}

				group, err := groupStore.FindGroup(ctx, filter)
				if err != nil {
					panic(err)
				}

				message.Group = &gchatmdl.Group{
					MongoId:  group.MongoId,
					Name:     group.Name,
					ImageUrl: group.ImageUrl,
				}

				// Extract sender from message.SenderId
				filter, err = common.GetIdFilter(message.SenderId)
				if err != nil {
					panic(err)
				}

				message.Sender, err = groupChatStore.FindUser(ctx, filter)
				if err != nil {
					panic(err)
				}

				// Send socket and notification to each member

				t := true
				f := false
				for _, userId := range group.Members {

					if userId == message.SenderId {
						message.IsMe = &t
					} else {
						message.IsMe = &f
					}
					err = appCtx.Socket().Send(userId, message)
					if err != nil {
						log.Err(err).Msg(err.Error())
					}
				}

			}(context.Background(), messageId)
		}
	}()
}
