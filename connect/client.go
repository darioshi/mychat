package connect

import (
	"context"
	"github.com/gorilla/websocket"
	"time"
)

type Client struct {
	ID              string             // 链接的唯一标识
	conn            *websocket.Conn    // 链接实体
	Ctx             context.Context    // 文本流
	CancelFunc      context.CancelFunc // 关闭函数
	lastRequestTime time.Time          // 上次服务端接收消息的时间
	retryTime       int                // 重试次数
}
