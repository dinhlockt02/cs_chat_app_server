package routeinternal

import (
	"cs_chat_app_server/components/appcontext"
	authmiddleware "cs_chat_app_server/middleware/authenticate"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitSocketRoute(g *gin.RouterGroup, appCtx appcontext.AppContext) {
	g.GET("/ws/chat", authmiddleware.Authentication(appCtx), func(c *gin.Context) {

		c.JSON(http.StatusMovedPermanently, gin.H{"data": "Move to /v1/friend/{{FRIEND_ID}}/chat/ws"})
		return

		//conn, _, _, err := ws.UpgradeHTTP(c.Request, c.Writer)
		//if err != nil {
		//	panic(common.ErrInternal(err))
		//}
		//u, _ := c.Get(common.CurrentUser)
		//requester, _ := u.(common.Requester)
		//err = appCtx.Socket().AddConn(requester.GetId(), conn)
		//if err != nil {
		//	panic(common.ErrInternal(err))
		//}
		//appCtx.Socket().Receive(conn, c, pchatskt.SendMessageHandler(appCtx))
	})

}
