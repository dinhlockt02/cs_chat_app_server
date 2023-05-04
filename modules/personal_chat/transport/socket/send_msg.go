package pchatskt

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	"cs_chat_app_server/components/socket"
	pchatbiz "cs_chat_app_server/modules/personal_chat/biz"
	pchatmdl "cs_chat_app_server/modules/personal_chat/model"
	pchatstore "cs_chat_app_server/modules/personal_chat/store"
	"encoding/json"
)

func SendMessageHandler(appCtx appcontext.AppContext) socket.SocketHandler {
	return func(c *socket.Context, data []byte) {
		defer func() {
			if err := recover(); err != nil {
				c.Response(err)
			}
		}()
		u, _ := c.GetContext().Get(common.CurrentUser)
		requester := u.(common.Requester)
		var item pchatmdl.PersonalChatItem
		err := json.Unmarshal(data, &item)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := pchatstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		biz := pchatbiz.NewSendMessageBiz(store, appCtx.Socket())
		item.Sender = requester.GetId()
		if err = biz.Send(context.Background(), &item); err != nil {
			panic(err)
		}
		c.Response(item)
		return
	}
}
