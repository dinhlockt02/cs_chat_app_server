package subscriber

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	groupmdl "cs_chat_app_server/modules/group/model"
	groupstore "cs_chat_app_server/modules/group/store"
	gchatstore "cs_chat_app_server/modules/group_chat/store"
	"errors"
	"github.com/rs/zerolog/log"
)

func UpdateGroupLatestMessageWhenNewGroupMessageReceived(appCtx appcontext.AppContext, ctx context.Context) {
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
					log.Error().
						Str("package", "subscriber.UpdateGroupLatestMessageWhenNewGroupMessageReceived").
						Err(err).Msg("can not create filter from messageId: " + messageId)
					return
				}

				message, err := groupChatStore.FindMessage(ctx, filter)

				if err != nil {
					log.Error().
						Str("package", "subscriber.UpdateGroupLatestMessageWhenNewGroupMessageReceived").
						Err(err).Msg("error while find message: " + messageId)
					return
				}

				if message == nil {
					err = errors.New("message not found")
					log.Error().
						Str("package", "subscriber.UpdateGroupLatestMessageWhenNewGroupMessageReceived").
						Err(err).Msg(err.Error())
					return
				}

				// Get sender
				filter, err = common.GetIdFilter(message.SenderId)
				if err != nil {
					log.Error().
						Str("package", "subscriber.UpdateGroupLatestMessageWhenNewGroupMessageReceived").
						Err(err).Msg("can not create filter from message.SenderId: " + message.SenderId)
					return
				}

				sender, err := groupChatStore.FindUser(ctx, filter)

				if err != nil {
					log.Error().
						Str("package", "subscriber.UpdateGroupLatestMessageWhenNewGroupMessageReceived").
						Err(err).Msg("error while find sender: " + message.SenderId)
					return
				}

				if sender == nil {
					err = errors.New("sender not found")
					log.Error().
						Str("package", "subscriber.UpdateGroupLatestMessageWhenNewGroupMessageReceived").
						Err(err).Msg(err.Error())
					return
				}

				// Update group from groupId
				filter, err = common.GetIdFilter(message.GroupId)
				if err != nil {
					log.Error().
						Str("package", "subscriber.UpdateGroupLatestMessageWhenNewGroupMessageReceived").
						Err(err).Msg("can not create filter from message.GroupId: " + message.GroupId)
					return
				}

				err = groupStore.UpdateLatestMessage(ctx, filter, &groupmdl.GroupMessage{
					Message:        message.Message,
					SenderId:       message.SenderId,
					SenderName:     sender.Name,
					MongoCreatedAt: message.MongoCreatedAt,
				})
				if err != nil {
					log.Error().
						Str("package", "subscriber.UpdateGroupLatestMessageWhenNewGroupMessageReceived").
						Err(err).Msg("error while update group: " + message.SenderId)
					return
				}

			}(context.Background(), messageId)
		}
	}()
}
