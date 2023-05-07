package socket

import (
	"cs_chat_app_server/common"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"net"
	"sync"
)

type wsSocket struct {
	conns map[string][]net.Conn
	mutex sync.RWMutex
}

func NewWSSocket() *wsSocket {
	return &wsSocket{
		conns: make(map[string][]net.Conn),
		mutex: sync.RWMutex{},
	}
}

func (w *wsSocket) AddConn(topic string, conn net.Conn) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	if _, ok := w.conns[topic]; ok {
		w.conns[topic] = append(w.conns[topic], conn)
	} else {
		w.conns[topic] = []net.Conn{conn}
	}
	return nil
}

func (w *wsSocket) RemoveConn(topic string, conn net.Conn) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	if _, ok := w.conns[topic]; ok {
		for i, _ := range w.conns[topic] {
			if w.conns[topic][i] == conn {
				w.conns[topic] = append(w.conns[topic][:i], w.conns[topic][i+1:]...)
			}
		}
	}
	return nil
}

func (w *wsSocket) Send(topic string, message interface{}) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	b, err := json.Marshal(message)
	if err != nil {
		return common.ErrInternal(err)
	}

	if _, ok := w.conns[topic]; !ok {
		return nil
	}
	for i, _ := range w.conns[topic] {
		err = wsutil.WriteServerMessage(w.conns[topic][i], ws.OpText, b)
		if err != nil {
			continue
		}
	}
	return nil
}

func (w *wsSocket) Receive(conn net.Conn, ginContext *gin.Context, handler SocketHandler) {
	go func() {
		for {
			msg, _, err := wsutil.ReadClientData(conn)
			if err != nil {
				break
			}
			handler(newContext(conn, ginContext), msg)
		}
	}()
}
