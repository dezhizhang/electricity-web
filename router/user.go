package router

import (
	"electricity-web/controller"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router gin.RouterGroup) {
	UserRouter := Router.Group("/api/v1/user")
	{
		UserRouter.GET("/list", controller.UserController.GetUserList)
	}

}
