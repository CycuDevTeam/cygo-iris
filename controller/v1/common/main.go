package common

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

type BasicResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(ctx iris.Context, data interface{}) {
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(BasicResponse{
		Code:    CodeSuccess,
		Message: ctx.Tr("SUCCESS"),
		Data:    data,
	})
}

// Client Error Response
// ----------------------------------------------------------------------------------
func ParamErrorResponse(ctx iris.Context, msg string) {
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(BasicResponse{Code: CodeParamError, Message: ctx.Tr(msg)})
	return
}

func FormErrorResponse(ctx iris.Context, err error) {
	msg := ""
	if errs, ok := err.(validator.ValidationErrors); ok {
		validationErrors := ""
		flag := true
		for _, validationErr := range errs {
			actualTag := ctx.Tr(validationErr.ActualTag())
			namespace := ctx.Tr(validationErr.Field())

			message := fmt.Sprintf("%v %v", namespace, actualTag)
			if flag {
				validationErrors = fmt.Sprintf("%v", message)
				flag = false
			} else {
				validationErrors = fmt.Sprintf("%v | %v", validationErrors, message)
			}
		}
		msg = validationErrors
	} else {
		msg = ctx.Tr("PARAM_ERROR")
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(BasicResponse{Code: CodeParamError, Message: msg})
}

// Server Error Response
// ----------------------------------------------------------------------------------
func DatabaseErrorResponse(ctx iris.Context) {
	ctx.StatusCode(iris.StatusInternalServerError)
	ctx.JSON(BasicResponse{Code: CodeDBError, Message: ctx.Tr("SERVER_ERROR")})
	return
}

func EncryptErrorResponse(ctx iris.Context) {
	ctx.StatusCode(iris.StatusInternalServerError)
	ctx.JSON(BasicResponse{Code: CodeEncryptError, Message: ctx.Tr("SERVER_ERROR")})
	return
}
