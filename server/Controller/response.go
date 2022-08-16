package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    MyCode      `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(message string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{200, message, data})
}

func Failed(message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{400, message, 0})
}

func ResponseErrorWithMsg(ctx *gin.Context, code MyCode, data interface{}) {
	res := &Response{
		Code:    code,
		Message: code.Msg(),
		Data:    data,
	}
	ctx.JSON(http.StatusOK, res)
}

func ResponseError(ctx *gin.Context, code MyCode) {
	res := &Response{
		Code:    code,
		Message: code.Msg(),
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, res)
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	res := &Response{
		Code:    CodeSuccess,
		Message: CodeSuccess.Msg(),
		Data:    data,
	}
	ctx.JSON(http.StatusOK, res)
}
