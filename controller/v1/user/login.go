package user

import (
	"cygo_iris/controller/v1/common"
	"cygo_iris/service"
	"github.com/kataras/iris/v12"
)

type LoginRequest struct {
	Account  string `json:"account" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponseData struct {
	Uid      string `json:"uid"`
	Username string `json:"username"`
	Role     int    `json:"role"`
}

func Login(ctx iris.Context) {
	var req LoginRequest
	if err := ctx.ReadJSON(&req); err != nil {
		common.FormErrorResponse(ctx, err)
		return
	}

	user := service.GetUserByAccount(req.Account)
	if !user.CheckPassword(req.Password) {
		common.ParamErrorResponse(ctx, "PASSWORD_ERROR")
		return
	}
	service.SetLoginSession(ctx, user.Uid)
	common.SuccessResponse(ctx, LoginResponseData{
		Uid:      user.Uid.Hex(),
		Username: user.Username,
		Role:     user.Role,
	})
}
