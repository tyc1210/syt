package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
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
	// 静态资源服务器
	router.StaticFS("/static/imgs", http.Dir("storage/imgs"))

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
	// 医院管理
	hospital := v1.NewHospital()
	{
		// 分页获取医院信息
		r.GET("/hosp/hospital/:page/:limit", hospital.GetHospital)
		// 获取医院等级分类
		r.GET("/hosp/level", hospital.GetLevel)
		// 获取医院地区分类
		r.GET("/hosp/region", hospital.GetRegion)
	}
	return router
}
