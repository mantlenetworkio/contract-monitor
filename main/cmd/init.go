package cmd

import (
	"fmt"
	"github.com/493labs/contract-monitor/config"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"path/filepath"
)

const (
	flagNameOfSystemConfigFilepath            = "sconf"
	flagNameShortHandOfSystemConfigFilepath   = "sc"
	flagNameOfContractConfigFilepath          = "cconf"
	flagNameShortHandOfContractConfigFilepath = "cc"
	flagNameOfBinaryDirPath                   = "dir"
	flagNameShortHandOfBinaryDirPath          = "d"
)

// initLocalConfig init local config
func initLocalConfig(cmd *cobra.Command) {
	if err := config.InitLocalConfig(cmd); err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func initContractConfig(cmd *cobra.Command) {
	if err := config.InitContractConfig(cmd); err != nil {
		fmt.Println(err)
		panic(err)
	}
}

// initFlagSet init flag set
func initFlagSet() *pflag.FlagSet {
	dir := filepath.Dir(".")
	flags := &pflag.FlagSet{}
	flags.StringVarP(&config.SystemConfigFilepath, flagNameOfSystemConfigFilepath, flagNameShortHandOfSystemConfigFilepath, config.SystemConfigFilepath, "specify system config file path, if not set, default use ./files/system_config.yml")
	flags.StringVarP(&config.ContractConfigFilePath, flagNameOfContractConfigFilepath, flagNameShortHandOfContractConfigFilepath, config.ContractConfigFilePath, "specify contract config file path, if not set, default use ./files/contract_config.yml")
	flags.StringVarP(&config.BinaryAbsDirPath, flagNameOfBinaryDirPath, flagNameShortHandOfBinaryDirPath, dir, "specify binary dir path, if not set, default use filepath.Dir(.)")
	return flags
}

// attachFlags
func attachFlags(cmd *cobra.Command, flagNames []string) {
	flags := initFlagSet()
	cmdFlags := cmd.Flags()
	for _, flagName := range flagNames {
		if flag := flags.Lookup(flagName); flag != nil {
			cmdFlags.AddFlag(flag)
		}
	}
}
