package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"syt/pkg/errcode"
)

type CommonResult struct {
	Ctx *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"message"`
}

func NewCommonResult(ctx *gin.Context) *CommonResult {
	return &CommonResult{ctx}
}

func (result CommonResult) Success(data interface{}) {
	result.Ctx.JSON(http.StatusOK, gin.H{
		"code":    errcode.Success.GetCode(),
		"data":    data,
		"message": errcode.Success.GetMsg(),
	})
}

func (result CommonResult) Error(c *errcode.Code) {
	result.Ctx.JSON(http.StatusOK, gin.H{
		"code":    c.GetCode(),
		"data":    nil,
		"message": c.GetMsg(),
	})
}
