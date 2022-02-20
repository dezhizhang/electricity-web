package initialize

import (
	"electricity-web/routers"
	"github.com/gin-gonic/gin"
)

// 初始化路由

func Routers() *gin.Engine  {
	Router := gin.Default()
	//用户路由
	ApiGroup := Router.Group("/api/v1")
	routers.InitUserRouter(ApiGroup)
	return Router
}