package socket

import (
	"github.com/gin-gonic/gin"
	"net"
)

type SocketHandler func(ctx *Context, data []byte)

type Socket interface {
	AddConn(topic string, conn net.Conn) error
	RemoveConn(topic string, conn net.Conn) error
	Send(topic string, message interface{}) error
	Receive(conn net.Conn, ginContext *gin.Context, handler SocketHandler)
}
