package socket

import (
	"cs_chat_app_server/common"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"net"
)

type Context struct {
	conn       net.Conn
	ginContext *gin.Context
}

func newContext(conn net.Conn, ginContext *gin.Context) *Context {
	return &Context{
		conn:       conn,
		ginContext: ginContext,
	}
}

func (ctx *Context) GetConn() net.Conn {
	return ctx.conn
}

func (ctx *Context) GetContext() *gin.Context {
	return ctx.ginContext
}

func (ctx *Context) Response(data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		_ = wsutil.WriteServerText(ctx.conn, []byte(common.ErrInvalidRequest(err).RootErr.Error()))
		return
	}
	err = wsutil.WriteServerMessage(ctx.conn, ws.OpText, b)
	if err != nil {
		_ = wsutil.WriteServerText(ctx.conn, []byte(common.ErrInvalidRequest(err).RootErr.Error()))
		return
	}

}
