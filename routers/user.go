package routers

import (
	"electricity-web/controller"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("/user")
	{
		UserRouter.GET("/list", controller.UserController.GetUserList)
	}

}
