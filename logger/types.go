package logger

// LogModuleConfig the log config of service
type LogModuleConfig struct {
	ModuleName   string `mapstructure:"module"`    // 归属模块
	LogLevel     string `mapstructure:"log_level"` // 日志等级
	FilePath     string `mapstructure:"file_path"` // 日志文件路径
	MaxAge       int    `mapstructure:"max_age"`   // 日志留存配置
	RotationTime int    `mapstructure:"rotation_time"`
	LogInConsole bool   `mapstructure:"log_in_console"` // 在标准输出中打印
	ShowColor    bool   `mapstructure:"show_color"`     // 显示颜色
}
