package config

import (
	"github.com/493labs/contract-monitor/logger"
)

var (
	SystemConfigFilepath   = "./files/system_config.yml"   // common config path
	SystemConfig           = &LocalConf{}                  // local config instance for global
)

type LocalConf struct {
	LogConfig      []*logger.LogModuleConfig `mapstructure:"log"`      // 日志配置
}