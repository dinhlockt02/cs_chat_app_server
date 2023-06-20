package routeinternal

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	authmiddleware "cs_chat_app_server/middleware/authenticate"
	friendgin "cs_chat_app_server/modules/friend/transport/gin"
	pchatgin "cs_chat_app_server/modules/personal_chat/transport/gin"
	pchatskt "cs_chat_app_server/modules/personal_chat/transport/socket"
	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
	"github.com/rs/zerolog/log"
)

func InitFriendRoute(g *gin.RouterGroup, appCtx appcontext.AppContext) {

	friend := g.Group("/friend", authmiddleware.Authentication(appCtx))
	{
		friend.GET("/", friendgin.ListFriend(appCtx))
		friendRequest := friend.Group("/request")
		{
			friendRequest.GET("/sent", friendgin.GetSentRequest(appCtx))
			friendRequest.GET("/received", friendgin.GetReceivedRequest(appCtx))
			friendRequest.POST("/:id", friendgin.SendRequest(appCtx))
			friendRequest.DELETE("/:id", friendgin.RecallRequest(appCtx))
			friendRequest.POST("/:id/accept", friendgin.AcceptRequest(appCtx))
			friendRequest.DELETE("/:id/reject", friendgin.RejectRequest(appCtx))
		}
		friend.DELETE("/:id", friendgin.Unfriend(appCtx))
		friend.PUT("/:id/block", friendgin.Block(appCtx))
		friend.PUT("/:id/unblock", friendgin.Unblock(appCtx))
		{
			friend.GET("/:id/chat", pchatgin.ListMessage(appCtx))
			friend.GET("/:id/chat/ws", friendWebsocketChatHandler(appCtx))
		}
	}
}

func friendWebsocketChatHandler(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Debug().Msg("New connection connected")
		id := c.Param("id")

		if _, err := common.ToObjectId(id); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		conn, _, _, err := ws.UpgradeHTTP(c.Request, c.Writer)
		if err != nil {
			panic(common.ErrInternal(err))
		}

		ctx := context.WithValue(context.Background(), common.CurrentFriendId, id)
		u, _ := c.Get(common.CurrentUser)
		ctx = context.WithValue(ctx, common.CurrentUser, u)
		requester, _ := u.(common.Requester)
		err = appCtx.Socket().AddConn(requester.GetId(), conn)
		if err != nil {
			panic(common.ErrInternal(err))
		}

		appCtx.Socket().Receive(conn, ctx, pchatskt.SendMessageHandler(appCtx))
	}
}
