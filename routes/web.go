package routes

import (
	"github.com/gin-gonic/gin"
	"goletter-go-im/app/controllers"
	"goletter-go-im/app/services/ws"
)

// 注册路由列表
func WebRouter(router *gin.RouterGroup) {
	router.GET("/", controllers.Home)
	router.GET("/chatIndex", controllers.PageChat)
	router.GET("/chat_main", controllers.PageChatMain)

	// 游客信息
	router.POST("/visitor_login", controllers.PostVisitorLogin)
	router.GET("/notice", controllers.GetNotice)
	router.GET("/messages", controllers.MessageIndex)

	// 客服信息
	router.GET("/kefuinfo", controllers.GetKefuInfo)

	//发送单条消息
	router.POST("/2/message", controllers.SendMessage)

	// 上传
	router.POST("/uploadimg", controllers.UploadImg)
	router.POST("/uploadfile", controllers.UploadFile)

	// 后台
	router.GET("/visitors_kefu_online", controllers.GetKefusVisitorOnlines)
	router.GET("/visitor", controllers.GetVisitor)
	router.GET("/visitors", controllers.GetVisitors)
	router.GET("/replys", controllers.GetReplys)
	router.GET("/ipblacks", controllers.GetIpblacksByKefuId)

	// socket
	router.GET("/chat_server", controllers.NewChatServer)
	router.GET("/ws_kefu", ws.NewKefuServer)
	router.GET("/ws_visitor", ws.NewVisitorServer)
	go ws.WsServerBackend()
}