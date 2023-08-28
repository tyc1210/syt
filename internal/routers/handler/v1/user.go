package v1

import (
	"github.com/gin-gonic/gin"
	"strings"
	"syt/internal/entry/request"
	"syt/pkg/app"
	"syt/pkg/errcode"
)

var (
	Token = "d3a528f0-4384-42e3-8209-5ab20e311ffb"
)

type User struct {
}

func NewUser() User {
	return User{}
}

// Login
// @Summary 登录
// @Tags 用户管理
// @Produce  json
// @param data body request.LoginRequest true "用户名密码"
// @Success 200 {object} app.Response "成功"
// @Router /api/login [post]
func (u User) Login(c *gin.Context) {
	result := app.NewCommonResult(c)
	var loginRequest request.LoginRequest
	err := c.ShouldBindJSON(&loginRequest)
	if err != nil {
		code := errcode.NewCode(errcode.InvalidParams.GetCode(), err.Error())
		result.Error(&code)
		return
	}
	if strings.Compare(loginRequest.UserName, "admin") == 0 && strings.Compare(loginRequest.PassWord, "123456") == 0 {
		result.Success(Token)
	} else {
		result.Error(&errcode.UserNameOrPwdError)
	}

}
