package controllers

import (
	"github.com/gin-gonic/gin"
	"goletter-go-im/app/models"
	"goletter-go-im/pkg/model"
	"goletter-go-im/pkg/utils"
)

func MessageIndex(c *gin.Context) {
	visitorId := c.Query("visitorId")
	messages := models.FindMessageByVisitorId(visitorId)
	result := make([]map[string]interface{}, 0)
	var visitor models.Visitor
	var kefu models.User

	for _, message := range messages {
		item := make(map[string]interface{})
		if visitor.Name == "" || kefu.Username == "" {
			kefu = models.FindUser(message.KefuId)
			visitor = models.FindVisitorByVistorId(message.VisitorId)
		}
		item["time"] = message.CreatedAt.Format("2006-01-02 15:04:05")
		item["content"] = message.Content
		item["mes_type"] = message.MesType
		item["visitor_name"] = visitor.Name
		item["visitor_avator"] = visitor.Avator
		item["kefu_name"] = kefu.Nickname
		item["kefu_avator"] = kefu.Avator
		result = append(result, item)

	}
	models.ReadMessageByVisitorId(visitorId)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": result,
	})
}

func MessageStore(c *gin.Context) {
	message := &models.Message{
		KefuId: "kefu2",
		VisitorId: "123456",
		Content: "测试",
		MesType: "visitor",
		Status: "unread",
	}
	model.Create(message)

	utils.ResponseSuccess(c, message)
	// utils.ResponseError(c, 2002, err)
	return
}

func MessageUpdate(c *gin.Context) {
	message := &models.Message{}
	model.First(message, 10)
	message.Content = "888889999999"
	model.UpdateOne(message)

	utils.ResponseSuccess(c, message)
	return
}

func MessageDelete(c *gin.Context) {
	message :=models.Message{}
	model.DeleteById(message, 10)

	utils.ResponseSuccess(c, message)
	return
}