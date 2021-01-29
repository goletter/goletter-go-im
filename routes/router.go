package routes

import (
	"github.com/gin-gonic/gin"
)

// 注册当前
func Register() *gin.Engine {
	router := gin.Default() // 获取路由实例

	//加载模板文件
	router.Static("/static", "./static")
	router.LoadHTMLGlob("views/*")

	// 定义web组
	web := router.Group("/")
	WebRouter(web)

	// 定义api组
	api := router.Group("/api")
	ApiRouter(api)

	return router // 返回路由
}
