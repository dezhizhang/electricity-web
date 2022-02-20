package controller

import (
	"context"
	"electricity-web/proto"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type UserController struct {
}

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg":  e.Message(),
					"code": 404,
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg":  "服务器出错",
					"code": 500,
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg":  "参数有误",
					"code": "405",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg":  "服务器出错",
					"code": 500,
				})

			}
		}
	}
	return
}

func (u *UserController) GetUserList(ctx *gin.Context) {
	ip := "127.0.0.1"
	port := 8000
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList]连接服务失败", "msg", err.Error())
		HandleGrpcErrorToHttp(err,ctx)
		return
	}

	userSrvClient := proto.NewUserClient(conn)
	rsp, err := userSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		Page:     1,
		PageSize: 10,
	})

	if err != nil {
		zap.S().Errorw("[GetUserList]查询失败")
		HandleGrpcErrorToHttp(err,ctx)
		return
	}

	result := make([]interface{},0)

	for _,value := range rsp.Data {
		data := make(map[string]interface{})
		data["id"] = value.Id
		data["name"] = value.NickName
		data["birthday"] = value.Birthday
		data["phone"] = value.Phone
		data["gender"] = value.Gender
		result = append(result,data)
	}

	ctx.JSON(http.StatusOK,result)

}
