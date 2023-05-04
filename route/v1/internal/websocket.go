package routeinternal

import (
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	authmiddleware "cs_chat_app_server/middleware/authenticate"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
	"github.com/rs/zerolog/log"
)

func InitSocketRoute(g *gin.RouterGroup, appCtx appcontext.AppContext) {
	g.GET("/ws", authmiddleware.Authentication(appCtx), func(c *gin.Context) {
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
		appCtx.Socket().Receive(conn, handler)
	})
}

func handler(data []byte) {
	type message struct {
		Test string `json:"test"`
	}
	var msg message
	err := json.Unmarshal(data, &msg)
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Info().Msg(msg.Test)
}
