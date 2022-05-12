package config

var (
	ContractConfig         = &ContractMonitors{}           // contract monitor config
	ContractConfigFilePath = "./files/contract_config.yml" // contract monitor config path
)

type ContractMonitors struct {
	CMonitors []ContractMonitor `mapstructure:"c_monitors"`
}

type ContractMonitor struct {
	Network string       `mapstructure:"network"`
	Name    string       `mapstructure:"name"`
	Address string       `mapstructure:"address"`
	ABI     string       `mapstructure:"abi"`
	Events  []EventConf  `mapstructure:"events"`
	Methods []MethodConf `mapstructure:"methods"`
}

type EventConf struct {
	Name       string `mapstructure:"name"`
	EventID    string `mapstructure:"event_id"`
	AlertLevel string `mapstructure:"alert_level"`
	AlertURL   string `mapstructure:"alert_url"`
}

type MethodConf struct {
	Name       string `mapstructure:"name"`
	MethodID   string `mapstructure:"method_id"`
	AlertType  string `mapstructure:"alert_type"`
	AlertLevel string `mapstructure:"alert_level"`
	AlertURL   string `mapstructure:"alert_url"`
}
