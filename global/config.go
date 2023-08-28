package global

import (
	"syt/internal/setting"
)

type Config struct {
	Server *setting.ServerProperties
	App    *setting.AppProperties
}

var (
	Cfg *Config
)
