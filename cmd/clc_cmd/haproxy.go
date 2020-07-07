package clc_cmd

import (
	"fmt"
	"github.com/Molsbee/jarvis/service"
	"github.com/spf13/cobra"
	"os"
)

func haProxy() *cobra.Command {
	ha := &cobra.Command{
		Use: "haproxy",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	ha.AddCommand(&cobra.Command{
		Use:     "stats",
		Example: "jarvis clc haproxy stats uc1",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("please provide a data center")
				os.Exit(1)
			}

			stats, err := service.GetHAProxyStatsPage(args[0])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			for _, stat := range stats {
				fmt.Println(stat)
			}
		},
	})

	return ha
}
