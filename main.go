package main

import (
	"electricity-web/initialize"
	"fmt"
	"go.uber.org/zap"
)

func main()  {
	port := 8080
	logger,_ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
	//初始化路由
	Router := initialize.Routers()

	zap.S().Debugf("启动服务器,端口:%d",port)
	//启动服务
	err := Router.Run(fmt.Sprintf(":%d",port))
	if err != nil {
		zap.S().Panic("启动服务失败",err.Error())
	}
}
