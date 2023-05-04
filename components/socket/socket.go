package socket

import "net"

type Socket interface {
	AddConn(userId string, conn net.Conn) error
	RemoveConn(userId string, conn net.Conn) error
	Send(userId string, message interface{}) error
	Receive(conn net.Conn, handler func(data []byte))
}
