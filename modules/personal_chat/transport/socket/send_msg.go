package pchatskt

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	"cs_chat_app_server/components/socket"
	"cs_chat_app_server/middleware"
	pchatbiz "cs_chat_app_server/modules/personal_chat/biz"
	pchatmdl "cs_chat_app_server/modules/personal_chat/model"
	pchatrepo "cs_chat_app_server/modules/personal_chat/repository"
	pchatstore "cs_chat_app_server/modules/personal_chat/store"
	"encoding/json"
)

func SendMessageHandler(appCtx appcontext.AppContext) socket.SocketHandler {
	return func(c *socket.Context, data []byte) {
		defer middleware.RecoverSocket(c)
		u, _ := c.GetContext().Get(common.CurrentUser)
		requester := u.(common.Requester)
		var item pchatmdl.PersonalChatItem
		err := json.Unmarshal(data, &item)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := pchatstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		repo := pchatrepo.NewCreateMessageRepo(store)
		biz := pchatbiz.NewSendMessageBiz(repo, appCtx.PubSub())

		item.SenderId = requester.GetId()
		receiverId, _ := c.GetContext().Get(common.CurrentFriendId)
		item.ReceiverId, _ = receiverId.(string)
		if err = biz.Send(context.Background(), &item); err != nil {
			panic(err)
		}
		//c.Response(map[string]interface{}{
		//	"data": item,
		//})
		return
	}
}
