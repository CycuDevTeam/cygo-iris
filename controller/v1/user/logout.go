package user

import (
	"cygo_iris/controller/v1/common"
	"cygo_iris/service"

	"github.com/kataras/iris/v12"
)

type LogoutRequest struct{}

type LogoutResponseData struct{}

func Logout(ctx iris.Context) {
	service.SetLogoutSession(ctx)
	common.SuccessResponse(ctx, LogoutResponseData{})
}
