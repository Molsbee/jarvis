package cmd

import (
	"fmt"
	"github.com/Molsbee/jarvis/service/clc"
	"github.com/spf13/cobra"
	"os"
	"regexp"
)

func clcCommands() *cobra.Command {
	clcCommand := &cobra.Command{
		Use:   "clc",
		Short: "Helpful commands for interacting with Century Link Cloud",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	find := &cobra.Command{
		Use:     "find",
		Aliases: []string{"search", "lookup"},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	find.AddCommand(&cobra.Command{
		Use:     "ip {{ ipAddress }}",
		Example: "jarvis clc find ip 10.121.12.15",
		Aliases: []string{"ipAddress", "ip-address"},
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				return
			}

			ipAddress := args[0]
			matched, err := regexp.Match(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`, []byte(ipAddress))
			if err != nil || !matched {
				fmt.Println("please provide a valid ip address")
				os.Exit(1)
			}

			address := clc.FindIPAddress(ipAddress)
			if address != nil {
				fmt.Println(address)
			} else {
				fmt.Printf("failed to find ip address %s", ipAddress)
			}
		},
	})

	clcCommand.AddCommand(find)
	return clcCommand
}
