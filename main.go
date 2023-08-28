package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strings"
	"syt/global"
	"syt/internal/routers"
	"syt/pkg/logger"
)

// @title 尚医通
// @version 1.0
// @description 尚医通项目后端（go语言版）
// @termsOfService https://github.com
func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:    ":" + global.Cfg.Server.HttpPort,
		Handler: router,
	}
	s.ListenAndServe()
}

func init() {
	InitConfig()
	InitLogger()
}

func InitLogger() {
	logger.InitLogger(global.Cfg.App)
}

func InitConfig() {
	viper := viper.New()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs/")
	viper.AutomaticEnv() // 启用自动读取环境变量的功能
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	err := viper.ReadInConfig() // 找到并加载配置文件
	if err != nil {             // 处理错误
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	if err := viper.Unmarshal(&global.Cfg); err != nil {
		log.Printf("unmarshal config file failed, %v", err)
	}
	log.Printf("server http port ===================> %s", global.Cfg.Server.HttpPort)
}
