package subscriber

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	pchatstore "cs_chat_app_server/modules/personal_chat/store"

	"github.com/rs/zerolog/log"
)

func NotifyUserWhenNewPersonalMessageReceived(appCtx appcontext.AppContext, ctx context.Context) {
	ch := appCtx.PubSub().Subscribe(ctx, common.TopicNewPersonalMessageCreated)

	personalChatStore := pchatstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
	go func() {
		for messageId := range ch {
			go func(ctx context.Context, messageId string) {
				defer common.Recovery()
				// Extract message from messageId
				filter, err := common.GetIdFilter(messageId)
				if err != nil {
					panic(err)
				}
				message, err := personalChatStore.FindMessage(ctx, filter)
				if err != nil {
					panic(err)
				}

				// Populate sender data into message

				filter, err = common.GetIdFilter(message.SenderId)
				if err != nil {
					panic(err)
				}

				message.Sender, err = personalChatStore.FindUser(ctx, filter)
				if err != nil {
					panic(err)
				}

				filter, err = common.GetIdFilter(message.SenderId)
				if err != nil {
					panic(err)
				}

				message.Receiver, err = personalChatStore.FindUser(ctx, filter)

				if err != nil {
					panic(err)
				}

				t := true
				f := false

				message.IsMe = &t

				err = appCtx.Socket().Send(message.SenderId, message)
				if err != nil {
					log.Err(err).Msg(err.Error())
				}
				message.IsMe = &f
				err = appCtx.Socket().Send(message.ReceiverId, message)
				if err != nil {
					log.Err(err).Msg(err.Error())
				}

			}(context.Background(), messageId)
		}
	}()
}
