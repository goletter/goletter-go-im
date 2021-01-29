package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"goletter-go-im/app/models"
	"goletter-go-im/app/services/ws"
	"goletter-go-im/pkg/model"
	"goletter-go-im/pkg/utils"
	"math/rand"
	"time"
)

func PostVisitorLogin(c *gin.Context) {
	toId := c.PostForm("to_id")
	avator := fmt.Sprintf("/static/images/%d.jpg", rand.Intn(14))
	id := c.PostForm("visitor_id")
	client_ip := c.ClientIP()
	city := "未识别地区"
	name := "匿名网友"

	if id == "" {
		id = utils.Uuid()
	}

	models.CreateVisitor(name, avator, c.ClientIP(), toId, id, city, client_ip)
	visitor :=models.FindVisitorByVistorId(id)

	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": visitor,
	})
}

func GetNotice(c *gin.Context) {
	kefuId := c.Query("kefu_id")
	user := models.FindUser(kefuId)

	result := make([]gin.H, 0)
	var welcomes []models.Welcome
	var condition []model.Condition

	model.GetAll(&welcomes, condition)
	for _, welcome := range welcomes {
		h := gin.H{
			"name":    user.Nickname,
			"avator":  user.Avator,
			"is_kefu": false,
			"content": welcome.Content,
			"time":    time.Now().Format("2006-01-02 15:04:05"),
		}
		result = append(result, h)
	}

	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": result,
	})
}

func SendMessage(c *gin.Context) {
	fromId := c.PostForm("from_id")
	toId := c.PostForm("to_id")
	content := c.PostForm("content")
	cType := c.PostForm("type")

	var kefuInfo models.User
	var vistorInfo models.Visitor
	if cType == "kefu" {
		kefuInfo = models.FindUser(fromId)
		vistorInfo = models.FindVisitorByVistorId(toId)
	} else if cType == "visitor" {
		vistorInfo = models.FindVisitorByVistorId(fromId)
		kefuInfo = models.FindUser(toId)
	}

	message := &models.Message{
		KefuId: kefuInfo.Username,
		VisitorId: vistorInfo.VisitorId,
		Content: content,
		MesType: cType,
		Status: "unread",
	}
	model.Create(message)

	var msg TypeMessage
	if cType == "kefu" {
		guest, ok := ws.ClientList[message.VisitorId]
		if guest != nil && ok {
			conn := guest.Conn

			msg = TypeMessage{
				Type: "message",
				Data: ws.ClientMessage{
					Name:    kefuInfo.Nickname,
					Avator:  kefuInfo.Avator,
					Id:      kefuInfo.Username,
					Time:    time.Now().Format("2006-01-02 15:04:05"),
					ToId:    message.VisitorId,
					Content: content,
				},
			}
			str, _ := json.Marshal(msg)
			conn.WriteMessage(websocket.TextMessage, str)
		}

		msg = TypeMessage{
			Type: "message",
			Data: ws.ClientMessage{
				Name:    kefuInfo.Nickname,
				Avator:  kefuInfo.Avator,
				Id:      vistorInfo.VisitorId,
				Time:    time.Now().Format("2006-01-02 15:04:05"),
				ToId:    vistorInfo.VisitorId,
				Content: content,
				IsKefu:  "yes",
			},
		}
		str2, _ := json.Marshal(msg)
		ws.OneKefuMessage(kefuInfo.Username, str2)
	}

	if cType == "visitor" {
		msg = TypeMessage{
			Type: "message",
			Data: ws.ClientMessage{
				Name:    kefuInfo.Nickname,
				Avator:  kefuInfo.Avator,
				Id:      message.VisitorId,
				Time:    time.Now().Format("2006-01-02 15:04:05"),
				ToId:    message.VisitorId,
				Content: content,
			},
		}
		str, _ := json.Marshal(msg)
		ws.OneKefuMessage(kefuInfo.Username, str)
	}

	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": msg,
	})
}
