package controllers

import (
	"github.com/gin-gonic/gin"
	"goletter-go-im/app/models"
	"goletter-go-im/pkg/model"
	"net/http"
)

func UserIndex(c *gin.Context) {
	pageInfo := model.PageInfo{
		Page: 1,
		PageSize: 10,
	}
	var users []models.User
	var condition []model.Condition
	   data :=model.Paginate(&users, pageInfo, condition)

	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": data,
	})
	// utils.ResponseSuccess(c, data)
	// utils.ResponseError(c, 2002, err)
}



func PageChat(c *gin.Context) {
	c.HTML(http.StatusOK, "chat_page.html", gin.H{})
}

func PageChatMain(c *gin.Context) {
	c.HTML(http.StatusOK, "chat_main.html", gin.H{})
}

func GetKefuInfo(c *gin.Context) {
	// kefuId, _ := c.Get("kefu_id")
	user := &models.User{}
	model.First(user, 1)
	info := make(map[string]interface{})
	info["name"] = user.Username
	info["id"] = user.Username
	info["avator"] = user.Avator
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": info,
	})
}

func GetKefusVisitorOnlines(c *gin.Context) {
	pageInfo := model.PageInfo{
		Page: 1,
		PageSize: 10,
	}
	var visitos []models.Visitor
	var condition []model.Condition
	data :=model.Paginate(&visitos, pageInfo, condition)

	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": data.Data,
	})
}

func GetVisitor(c *gin.Context) {
	visitorId := c.Query("visitorId")
	vistor := models.FindVisitorByVistorId(visitorId)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": vistor,
	})
}

func GetVisitors(c *gin.Context) {
	pageInfo := model.PageInfo{
		Page: 1,
		PageSize: 10,
	}
	var visitos []models.Visitor
	var condition []model.Condition
	data :=model.Paginate(&visitos, pageInfo, condition)

	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": data,
	})
}

func GetReplys(c *gin.Context) {
	// kefuId, _ := c.Get("kefu_name")
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": "",
	})
}

func GetIpblacksByKefuId(c *gin.Context) {
	// kefuId, _ := c.Get("kefu_name")
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": "",
	})
}