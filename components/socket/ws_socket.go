package socket

import (
	"cs_chat_app_server/common"
	"encoding/json"
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

func (w *wsSocket) AddConn(userId string, conn net.Conn) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	if _, ok := w.conns[userId]; ok {
		w.conns[userId] = append(w.conns[userId], conn)
	} else {
		w.conns[userId] = []net.Conn{conn}
	}
	return nil
}

func (w *wsSocket) RemoveConn(userId string, conn net.Conn) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	if _, ok := w.conns[userId]; ok {
		for i, _ := range w.conns[userId] {
			if w.conns[userId][i] == conn {
				w.conns[userId] = append(w.conns[userId][:i], w.conns[userId][i+1:]...)
			}
		}
	}
	return nil
}

func (w *wsSocket) Send(userId string, message interface{}) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	b, err := json.Marshal(message)
	if err != nil {
		return common.ErrInternal(err)
	}

	if _, ok := w.conns[userId]; !ok {
		return nil
	}
	for i, _ := range w.conns[userId] {
		err = wsutil.WriteServerMessage(w.conns[userId][i], ws.OpBinary, b)
		if err != nil {
			continue
		}
	}
	return nil
}

func (w *wsSocket) Receive(conn net.Conn, handler func(data []byte)) {
	go func() {
		for {
			msg, _, err := wsutil.ReadClientData(conn)
			if err != nil {
				break
			}
			handler(msg)
		}
	}()
}
