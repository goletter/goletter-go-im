package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

type vistor struct {
	conn   *websocket.Conn
	name   string
	id     string
	avator string
	to_id  string
}

var kefuList = make(map[string][]*websocket.Conn)

type Message struct {
	conn        *websocket.Conn
	c           *gin.Context
	content     []byte
	messageType int
}

type TypeMessage struct {
	Type interface{} `json:"type"`
	Data interface{} `json:"data"`
}

var upgrader = websocket.Upgrader{}

//定时检测客户端是否在线
func init() {
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
}

func NewChatServer(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	fmt.Println(conn)

	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	for {
		// 接受消息
	}
}