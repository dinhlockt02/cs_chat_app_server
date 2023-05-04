package socket

import (
	"github.com/gin-gonic/gin"
	"net"
)

type SocketHandler func(ctx *Context, data []byte)

type Socket interface {
	AddConn(userId string, conn net.Conn) error
	RemoveConn(userId string, conn net.Conn) error
	Send(userId string, message interface{}) error
	Receive(conn net.Conn, ginContext *gin.Context, handler SocketHandler)
}
