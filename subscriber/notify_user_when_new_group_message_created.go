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
					log.Error().Str("package", "subscriber.NotifyUserWhenNewGroupMessageReceived").
						Err(err).Msg("cant create message id filter: " + messageId)
					return
				}

				message, err := groupChatStore.FindMessage(ctx, filter)

				if err != nil {
					log.Error().Str("package", "subscriber.NotifyUserWhenNewGroupMessageReceived").
						Err(err).Msg("error while find message: " + messageId)
					return
				}

				if message == nil {
					log.Error().Str("package", "subscriber.NotifyUserWhenNewGroupMessageReceived").
						Err(err).Msg("message not found: " + messageId)
					return
				}

				// Extract group from groupId
				filter, err = common.GetIdFilter(message.GroupId)
				if err != nil {
					log.Error().Str("package", "subscriber.NotifyUserWhenNewGroupMessageReceived").
						Err(err).Msg("cant create group id filter: " + message.GroupId)
					return
				}

				group, err := groupStore.FindGroup(ctx, filter)
				if err != nil {
					log.Error().Str("package", "subscriber.NotifyUserWhenNewGroupMessageReceived").
						Err(err).Msg("error while find group: " + messageId)
					return
				}

				if group == nil {
					log.Error().Str("package", "subscriber.NotifyUserWhenNewGroupMessageReceived").
						Err(err).Msg("message not found: " + messageId)
					return
				}

				message.Group = &gchatmdl.Group{
					MongoId:  group.MongoId,
					Name:     group.Name,
					ImageUrl: group.ImageUrl,
					Type:     group.Type,
				}

				// Extract sender from message.SenderId
				filter, err = common.GetIdFilter(message.SenderId)
				if err != nil {
					log.Error().Str("package", "subscriber.NotifyUserWhenNewGroupMessageReceived").
						Err(err).Msg("cant create sender id filter: " + message.SenderId)
					return
				}

				message.Sender, err = groupChatStore.FindUser(ctx, filter)
				if err != nil {
					log.Error().Str("package", "subscriber.NotifyUserWhenNewGroupMessageReceived").
						Err(err).Msg("error while find sender: " + message.SenderId)
					return
				}

				// Send socket and notification to each member

				t := true
				f := false
				for _, member := range group.Members {

					if member.Id == message.SenderId {
						message.IsMe = &t
					} else {
						message.IsMe = &f
					}
					err = appCtx.Socket().Send(member.Id, message)
					if err != nil {
						log.Err(err).
							Str("package", "subscriber.NotifyUserWhenNewGroupMessageReceived").
							Msg(err.Error())
					}
				}

			}(context.Background(), messageId)
		}
	}()
}
