package gchatskt

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	"cs_chat_app_server/components/socket"
	"cs_chat_app_server/middleware"
	gchatbiz "cs_chat_app_server/modules/group_chat/biz"
	gchatmdl "cs_chat_app_server/modules/group_chat/model"
	gchatrepo "cs_chat_app_server/modules/group_chat/repository"
	gchatstore "cs_chat_app_server/modules/group_chat/store"
	"encoding/json"
	"github.com/rs/zerolog/log"
)

func SendMessageHandler(appCtx appcontext.AppContext) socket.SocketHandler {
	return func(c *socket.Context, data []byte) {
		defer middleware.RecoverSocket(c)
		u := c.GetContext().Value(common.CurrentUser)
		requester := u.(common.Requester)
		log.Debug().
			Str("package", "gchatskt.SendMessageHandler").
			Str("current user", requester.GetId()).
			Send()
		var item gchatmdl.GroupChatItem
		err := json.Unmarshal(data, &item)
		if err != nil {
			log.Debug().
				Err(err).
				Str("package", "gchatskt.SendMessageHandler").
				Msg("error while unmarshal message")
			panic(common.ErrInvalidRequest(err))
		}

		store := gchatstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		repo := gchatrepo.NewGroupChatRepository(store)
		biz := gchatbiz.NewSendMessageBiz(repo, appCtx.PubSub())

		item.SenderId = requester.GetId()
		receiverId := c.GetContext().Value(common.CurrentGroupId)
		log.Debug().
			Err(err).
			Str("package", "gchatskt.SendMessageHandler").
			Msg("current group id: " + receiverId.(string))
		item.GroupId, _ = receiverId.(string)
		if err = biz.Send(context.Background(), &item); err != nil {
			log.Debug().
				Err(err).
				Str("package", "gchatskt.SendMessageHandler").
				Msg("error while biz.Send")
			panic(err)
		}
		return
	}
}
