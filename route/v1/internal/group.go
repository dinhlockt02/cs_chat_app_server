package routeinternal

import (
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	authmiddleware "cs_chat_app_server/middleware/authenticate"
	groupgin "cs_chat_app_server/modules/group/transport/gin"
	gchatgin "cs_chat_app_server/modules/group_chat/transport/gin"
	gchatskt "cs_chat_app_server/modules/group_chat/transport/socket"
	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
)

func InitGroupRoute(g *gin.RouterGroup, appCtx appcontext.AppContext) {

	group := g.Group("/group", authmiddleware.Authentication(appCtx))
	{
		groupRequest := group.Group("/request")
		{
			groupRequest.GET("/sent", groupgin.GetSentRequest(appCtx))
			groupRequest.GET("/received", groupgin.GetReceiveRequest(appCtx))

			groupRequest.POST("/:groupId/accept", groupgin.AcceptRequest(appCtx))
			groupRequest.DELETE("/:groupId/reject", groupgin.RejectRequest(appCtx))

			groupRequest.POST("/:groupId/:friendId", groupgin.SendGroupRequest(appCtx))
			groupRequest.DELETE("/:groupId/:friendId", groupgin.RecallRequest(appCtx))
		}
		group.POST("", groupgin.CreateGroup(appCtx))
		group.GET("", groupgin.ListGroup(appCtx))
		group.PUT("/:groupId", groupgin.UpdateGroup(appCtx))

		{
			group.GET("/:groupId/chat", gchatgin.ListMessage(appCtx))
			group.GET("/:groupId/chat/ws", groupWebsocketChatHandler(appCtx))
		}

	}
}

func groupWebsocketChatHandler(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("groupId")
		if _, err := common.ToObjectId(id); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		conn, _, _, err := ws.UpgradeHTTP(c.Request, c.Writer)
		if err != nil {
			panic(common.ErrInternal(err))
		}

		c.Set(common.CurrentGroupId, id)

		u, _ := c.Get(common.CurrentUser)
		requester, _ := u.(common.Requester)
		err = appCtx.Socket().AddConn(requester.GetId(), conn)
		if err != nil {
			panic(common.ErrInternal(err))
		}

		appCtx.Socket().Receive(conn, c, gchatskt.SendMessageHandler(appCtx))
	}
}
