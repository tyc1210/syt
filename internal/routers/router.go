package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "syt/docs"
	"syt/internal/middleware"
	v1 "syt/internal/routers/handler/v1"
	"syt/pkg/logger"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.Cors())
	router.Use(logger.NewLogger())
	// 接口文档
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r := router.Group("/api")

	sms := v1.NewSms()
	// 短信验证码
	{
		r.GET("/sms/send/:phone", sms.GetCode)
	}
	// 用户管理
	user := v1.NewUser()
	{
		r.POST("/login", user.Login)
	}
	return router
}
