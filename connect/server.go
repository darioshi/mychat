package connect

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"mychat/proto"
	"time"
)

type Server struct {
	Options ServerOptions
}

type ServerOptions struct {
	WriteWait       time.Duration
	PongWait        time.Duration
	PingPeriod      time.Duration
	MaxMessageSize  int64
	ReadBufferSize  int
	WriteBufferSize int
	BroadcastSize   int
}

func NewServer(options ServerOptions) *Server {
	s := new(Server)
	s.Options = options
	return s
}

func (server *Server) writePump(ch *Channel) {

}

func (server *Server) readPump(ch *Channel) {
	defer func() {
		logrus.Infof("start exec disConnect ...")
		ch.conn.Close()
	}()

	ch.conn.SetReadLimit(server.Options.MaxMessageSize)
	ch.conn.SetReadDeadline(time.Now().Add(server.Options.PongWait))
	ch.conn.SetPongHandler(func(string) error {
		ch.conn.SetReadDeadline(time.Now().Add(server.Options.PongWait))
		return nil
	})

	for {
		_, message, err := ch.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logrus.Errorf("readPump ReadMessage err:%s", err.Error())
				return
			}
		}

		if message == nil {
			return
		}

		var connReq *proto.ConnectRequest
		logrus.Infof("get a message :%s", message)
		if err := json.Unmarshal([]byte(message), connReq); err != nil {
			logrus.Errorf("message struct %+v", connReq)
		}
	}
}
