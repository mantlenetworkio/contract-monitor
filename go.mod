module github.com/493labs/contract-monitor

go 1.16

require (
	github.com/ethereum/go-ethereum v1.10.17
	github.com/sirupsen/logrus v1.8.1
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	github.com/spf13/viper v1.11.0
)

replace (
	github.com/493labs/contract-monitor/types => ./types
	github.com/493labs/contract-monitor/config => ./config
	github.com/493labs/contract-monitor/logger => ./logger
)
