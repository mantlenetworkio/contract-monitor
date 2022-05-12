package main

import (
	"fmt"
	"github.com/493labs/contract-monitor/main/cmd"
	"github.com/spf13/cobra"
)

func main() {
	mainCmd := &cobra.Command{Use: "start"}
	mainCmd.AddCommand(cmd.StartCMD())

	err := mainCmd.Execute()
	if err != nil {
		_ = fmt.Errorf("monitor start error, %v", err)
	}
}
