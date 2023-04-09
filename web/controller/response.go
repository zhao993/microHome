package controller

import (
	"github.com/gin-gonic/gin"
	"microHome/web/utils"
)

type Response struct {
	Code utils.MyCode `json:"errno"`
	Msg  string       `json:"errmsg"`
	Data interface{}  `json:"data"`
}

func ResponseError(c *gin.Context, code utils.MyCode) {
	c.JSON(200, &Response{
		Code: code,
		Msg:  code.RecodeText(),
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(200, &Response{
		Code: utils.RecodeOk,
		Msg:  "成功",
		Data: data,
	})
}
