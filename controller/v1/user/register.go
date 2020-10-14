package user

import (
	"cygo_iris/controller/v1/common"
	"cygo_iris/repository"
	"cygo_iris/service"

	"strings"
	"time"

	"github.com/kataras/iris/v12"
	"gopkg.in/mgo.v2/bson"
)

type RegisterRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Username        string `json:"username" validate:"required,min=3,max=20,is_username"`
	Password        string `json:"password" validate:"required,min=8,max=20"`
	PasswordConfirm string `json:"passwordConfirm" validate:"eqfield=Password"`
}

type RegisterResponseData struct{}

func Register(ctx iris.Context) {
	var req RegisterRequest
	if err := ctx.ReadJSON(&req); err != nil {
		common.FormErrorResponse(ctx, err)
		return
	}

	if req.Username == "root" || req.Username == "admin" {
		common.ParamErrorResponse(ctx, "USERNAME_IN_USED")
		return
	}

	if exist, err := service.CheckUserExistByUsername(req.Username); err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	} else if exist {
		common.ParamErrorResponse(ctx, "USERNAME_IN_USED")
		return
	}

	if exist, err := service.CheckUserExistByEmail(req.Email); err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	} else if exist {
		common.ParamErrorResponse(ctx, "EMAIL_IN_USED")
		return
	}

	// Prepare user
	user := repository.User{
		Uid:         bson.NewObjectId(),
		Email:       strings.ToLower(req.Email),
		Verified:    false,
		Username:    strings.ToLower(req.Username),
		CreatedTime: time.Now().Unix(),
		LastLogin:   time.Now().Unix(),
		Role:        repository.NormalUser,
		IsBanned:    false,
	}
	if user.SetPassword(req.Password) != nil {
		common.EncryptErrorResponse(ctx)
		return
	}
	if err := service.InsertUser(&user); err != nil {
		common.DatabaseErrorResponse(ctx)
		return
	}
	common.SuccessResponse(ctx, RegisterResponseData{})
}
