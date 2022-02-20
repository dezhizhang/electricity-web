package main

import (
	"electricity-web/routers"
	"github.com/gin-gonic/gin"
)

func main()  {
	router := gin.Default()
	//用户路由
	ApiGroup := router.Group("/api/v1")
	routers.InitUserRouter(ApiGroup)
}
