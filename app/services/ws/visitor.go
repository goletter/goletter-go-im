package ws

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"goletter-go-im/app/models"
	"goletter-go-im/pkg/model"
	"log"
)

func NewVisitorServer(c *gin.Context) {
	//go kefuServerBackend()
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	vistor := &models.Visitor{}
	model.First(vistor, 1)

	user := &User{
		Conn:   conn,
		Name:   vistor.Name,
		Avator: vistor.Avator,
		Id:     vistor.VisitorId,
		To_id:  vistor.ToId,
	}
	AddVisitorToList(user)

	for {
		//接受消息
		var receive []byte
		messageType, receive, err := conn.ReadMessage()
		if err != nil {
			for _, visitor := range ClientList {
				if visitor.Conn == conn {
					log.Println("删除用户", visitor.Id)
					delete(ClientList, visitor.Id)
					VisitorOffline(visitor.To_id, visitor.Id, visitor.Name)
				}
			}
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

func AddVisitorToList(user *User) {
	//用户id对应的连接
	ClientList[user.Id] = user
	userInfo := make(map[string]string)
	userInfo["uid"] = user.Id
	userInfo["name"] = user.Name
	userInfo["avator"] = user.Avator
	userInfo["last_message"] = "11111"
	if userInfo["last_message"] == "" {
		userInfo["last_message"] = "新访客"
	}
	msg := TypeMessage{
		Type: "userOnline",
		Data: userInfo,
	}
	str, _ := json.Marshal(msg)
	//新版
	OneKefuMessage(user.To_id, str)
}

func VisitorOffline(kefuId string, visitorId string, visitorName string) {

	models.UpdateVisitorStatus(visitorId, 0)
	userInfo := make(map[string]string)
	userInfo["uid"] = visitorId
	userInfo["name"] = visitorName
	msg := TypeMessage{
		Type: "userOffline",
		Data: userInfo,
	}
	str, _ := json.Marshal(msg)
	//新版
	OneKefuMessage(kefuId, str)
}