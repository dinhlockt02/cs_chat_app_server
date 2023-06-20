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
)

func SendMessageHandler(appCtx appcontext.AppContext) socket.SocketHandler {
	return func(c *socket.Context, data []byte) {
		defer middleware.RecoverSocket(c)
		u := c.GetContext().Value(common.CurrentUser)
		requester := u.(common.Requester)
		var item gchatmdl.GroupChatItem
		err := json.Unmarshal(data, &item)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := gchatstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		repo := gchatrepo.NewGroupChatRepository(store)
		biz := gchatbiz.NewSendMessageBiz(repo, appCtx.PubSub())

		item.SenderId = requester.GetId()
		receiverId := c.GetContext().Value(common.CurrentGroupId)
		item.GroupId, _ = receiverId.(string)
		if err = biz.Send(context.Background(), &item); err != nil {
			panic(err)
		}
		return
	}
}
