package ws

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"goletter-go-im/app/models"
	"goletter-go-im/pkg/model"
	"log"
)


func NewKefuServer(c *gin.Context) {
	kefuInfo := &models.User{}
	model.First(kefuInfo, 1)

	//go kefuServerBackend()
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	//获取GET参数,创建WS
	var kefu User
	kefu.Id = kefuInfo.Username
	kefu.Name = kefuInfo.Nickname
	kefu.Avator = kefuInfo.Avator
	kefu.Conn = conn
	AddKefuToList(&kefu)

	for {
		//接受消息
		var receive []byte
		messageType, receive, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		message <- &Message{
			conn:        conn,
			content:     receive,
			context:     c,
			messageType: messageType,
		}
	}
}

//给指定客服发消息
func OneKefuMessage(toId string, str []byte) {
	//新版
	mKefuConns := KefuList[toId]
	if mKefuConns != nil {
		for _, kefu := range mKefuConns {
			kefu.Conn.WriteMessage(websocket.TextMessage, str)
		}
	}
}

func AddKefuToList(kefu *User) {
	var newKefuConns = []*User{kefu}
	kefuConns := KefuList[kefu.Id]

	if kefuConns != nil {
		for _, otherKefu := range kefuConns {
			msg := TypeMessage{
				Type: "many pong",
			}
			str, _ := json.Marshal(msg)
			err := otherKefu.Conn.WriteMessage(websocket.TextMessage, str)
			if err == nil {
				newKefuConns = append(newKefuConns, otherKefu)
			}
		}
	}
	KefuList[kefu.Id] = newKefuConns
}