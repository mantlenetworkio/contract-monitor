package config

import (
	"github.com/493labs/contract-monitor/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"path/filepath"
)

var (
	flagSets               = make([]*pflag.FlagSet, 0)     // flag set
	BinaryAbsDirPath       = ""                            // default release path
)

// InitLocalConfig init local config
func InitLocalConfig(cmd *cobra.Command) error {
	// 1. init config
	config, err := initLocal(cmd)
	if err != nil {
		return err
	}
	// 2. set log config
	logger.InitLogConfig(config.LogConfig)
	// 3. set global config and export
	SystemConfig = config
	return nil
}

func initLocal(cmd *cobra.Command) (*LocalConf, error){
	cmViper := viper.New()
	// 1. load the path of the config files
	ymlFile := SystemConfigFilepath
	if !filepath.IsAbs(ymlFile) {
		// 获取绝对路径
		ymlFile = FinalCfgPath(ymlFile)
		SystemConfigFilepath = ymlFile
	}
	// 2. load the config file
	cmViper.SetConfigFile(ymlFile)
	if err := cmViper.ReadInConfig(); err != nil {
		return nil, err
	}
	for _, command := range cmd.Commands() {
		flagSets = append(flagSets, command.PersistentFlags())
		err := cmViper.BindPFlags(command.PersistentFlags())
		if err != nil {
			return nil, err
		}
	}
	// 3. create new SystemConfig instance
	config := &LocalConf{}
	if err := cmViper.Unmarshal(config); err != nil {
		return nil, err
	}
	return config, nil
}

// InitContractConfig init monitor contract configs
func InitContractConfig(cmd *cobra.Command) error {
	// 1. init config
	config, err := initContract(cmd)
	if err != nil {
		return err
	}
	// 2. set global config and export
	ContractConfig = config
	return nil
}

func initContract(cmd *cobra.Command) (*ContractMonitors, error) {
	cmViper := viper.New()
	// 1. load the path of the config files
	ymlFile := ContractConfigFilePath
	if !filepath.IsAbs(ymlFile) {
		// 获取绝对路径
		ymlFile = FinalCfgPath(ymlFile)
		ContractConfigFilePath = ymlFile
	}
	// 2. load the config file
	cmViper.SetConfigFile(ymlFile)
	if err := cmViper.ReadInConfig(); err != nil {
		return nil, err
	}
	for _, command := range cmd.Commands() {
		flagSets = append(flagSets, command.PersistentFlags())
		err := cmViper.BindPFlags(command.PersistentFlags())
		if err != nil {
			return nil, err
		}
	}
	// 3. create new SystemConfig instance
	config := &ContractMonitors{}
	if err := cmViper.Unmarshal(config); err != nil {
		return nil, err
	}
	return config, nil
}

// FinalCfgPath check config and return absolute path
func FinalCfgPath(innerCfgPath string) string {
	var finalCfgPath string
	if filepath.IsAbs(innerCfgPath) {
		finalCfgPath = innerCfgPath
	} else {
		finalCfgPath = filepath.Join(BinaryAbsDirPath, innerCfgPath)
	}
	return finalCfgPath
}
