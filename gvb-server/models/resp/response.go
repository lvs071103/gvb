package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct{
	Code int `json:"code"`
	Data any `json:"data"`
	Msg string `json:"msg"`
}

const (
	Success = 0
	Err = 7	
)

func Result(code int, data any, msg string, c *gin.Context)  {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg: msg,
	})
}

func Ok(data any, msg string, c *gin.Context)  {
	Result(Success, data, msg, c)
}

func OkwithData(data any, c *gin.Context)  {
	Result(Success, data, "成功", c)
}

func OkwithMsg(msg string, c *gin.Context)  {
	Result(Success, map[string]any{}, msg, c)
}

func Failed(data any, msg string, c *gin.Context)  {
	Result(Err, data, msg, c)
}

func FailedwithMsg(msg string, c *gin.Context)  {
	Result(Err, map[string]any{}, msg, c)
}

func FailedWithCode(code ErrorCode, c *gin.Context)  {
	msg, ok := ErrorMap[ErrorCode(code)]
	if ok {
		Result(int(code), map[string]any{}, msg, c)
		return
	}
	Result(Err, map[string]any{}, "未知错误", c)
}