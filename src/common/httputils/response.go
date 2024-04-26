package httpUtil

import (
	"lookingforpartner/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code common.RespCode `json:"code"`
	Msg  interface{}     `json:"msg"`
	Data interface{}     `json:"data,omitempty"`
}

func ResponseErrorWithMsg(c *gin.Context, code common.RespCode, httpCode int, msg interface{}) {
	respData := &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	c.JSON(httpCode, respData)
}

func ResponseError(c *gin.Context, code common.RespCode, httpCode int) {
	respData := &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	c.JSON(httpCode, respData)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	respData := &ResponseData{
		Code: common.CodeSuccess,
		Msg:  common.CodeSuccess.Msg(),
		Data: data,
	}
	c.JSON(http.StatusOK, respData)
}
