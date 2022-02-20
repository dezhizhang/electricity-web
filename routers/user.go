package routers

import (
	"electricity-web/controller"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("/user")
	//用户控制器
	UserController := controller.UserController{}
	{
		UserRouter.GET("/list",UserController.GetUserList)
	}

}
