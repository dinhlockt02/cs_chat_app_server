package routeinternal

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	authmiddleware "cs_chat_app_server/middleware/authenticate"
	gchatskt "cs_chat_app_server/modules/group_chat/transport/socket"
	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
)

func InitSocketRoute(g *gin.RouterGroup, appCtx appcontext.AppContext) {
	g.GET("/ws", authmiddleware.Authentication(appCtx), websocketChatHandler(appCtx))

}

func websocketChatHandler(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, _, _, err := ws.UpgradeHTTP(c.Request, c.Writer)
		if err != nil {
			panic(common.ErrInternal(err))
		}
		u, _ := c.Get(common.CurrentUser)

		requester, _ := u.(common.Requester)
		err = appCtx.Socket().AddConn(requester.GetId(), conn)
		if err != nil {
			panic(common.ErrInternal(err))
		}

		appCtx.Socket().Receive(conn, context.Background(), gchatskt.SendMessageHandler(appCtx))
	}
}
