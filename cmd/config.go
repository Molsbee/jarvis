package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/Molsbee/jarvis/config"
	"github.com/spf13/cobra"
)

var showConfig = &cobra.Command{
	Use:   "show-config",
	Short: "Inspect the config file being sourced by the program",
	Run: func(cmd *cobra.Command, args []string) {
		j, _ := json.MarshalIndent(config.UserConfig, "", "  ")
		fmt.Println(string(j))
	},
}
