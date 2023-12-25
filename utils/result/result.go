package result

import (
	"gin-template/utils/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IError struct {
	Code     int
	ErrorMsg string
}

func NewIError(code int, err error) *IError {
	return &IError{
		Code:     code,
		ErrorMsg: err.Error(),
	}
}

func (e *IError) Error() string {
	return e.ErrorMsg
}

type Response struct {
	Code int         `json:"code"` // 错误码
	Msg  string      `json:"msg"`  // 错误描述
	Data interface{} `json:"data"` // 返回数据
}

// SetMsg 自定义响应信息
func (res *Response) SetMsg(message string) *Response {
	return &Response{
		Code: res.Code,
		Msg:  message,
		Data: res.Data,
	}
}

// SetData 追加响应数据
func (res *Response) SetData(data interface{}) *Response {
	return &Response{
		Code: res.Code,
		Msg:  res.Msg,
		Data: data,
	}
}

// SetCode  追加响应code
func (res *Response) SetCode(code int) *Response {
	return &Response{
		Code: code,
		Msg:  res.Msg,
		Data: res.Data,
	}
}

func (res *Response) GetReqMsg(req interface{}, err error) *Response {
	return &Response{
		Code: res.Code,
		Msg:  core.GetErrorMsg(req, err),
		Data: res.Data,
	}
}

func (res *Response) ToJson(c *gin.Context) {
	ret := Response{
		Code: res.Code,
		Msg:  res.Msg,
		Data: res.Data,
	}
	c.JSON(http.StatusOK, ret)
}

func Resp(code int, msg string) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}
