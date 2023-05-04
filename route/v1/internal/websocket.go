package routeinternal

import (
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	authmiddleware "cs_chat_app_server/middleware/authenticate"
	pchatskt "cs_chat_app_server/modules/personal_chat/transport/socket"
	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
)

func InitSocketRoute(g *gin.RouterGroup, appCtx appcontext.AppContext) {
	g.GET("/ws/chat", authmiddleware.Authentication(appCtx), func(c *gin.Context) {
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
		appCtx.Socket().Receive(conn, c, pchatskt.SendMessageHandler(appCtx))
	})
}
