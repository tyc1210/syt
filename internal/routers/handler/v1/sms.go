package v1

import (
	"github.com/gin-gonic/gin"
	"regexp"
	"syt/pkg/app"
	"syt/pkg/errcode"
)

type Sms struct {
}

func NewSms() Sms {
	return Sms{}
}

// GetCode
// @Summary 获取验证码
// @Tags 验证码
// @Produce  json
// @Param phone path string true "手机号"
// @Router /api/sms/send/{phone} [get]
// @Success 200 {object} app.Response "成功"
func (s Sms) GetCode(c *gin.Context) {
	phone := c.Param("phone")
	result := app.NewCommonResult(c)
	if validatePhoneNumber(phone) {
		result.Success("0917")
	} else {
		result.Error(&errcode.SmsPhoneError)
	}
}

func validatePhoneNumber(phoneNumber string) bool {
	// 定义手机号的正则表达式模式
	pattern := `^1[3456789]\d{9}$`

	// 编译正则表达式
	regex := regexp.MustCompile(pattern)

	// 使用正则表达式匹配手机号
	return regex.MatchString(phoneNumber)
}
