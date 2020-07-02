package clc_cmd

import (
	"fmt"
	"github.com/Molsbee/jarvis/service/clc"
	"github.com/spf13/cobra"
	"os"
	"regexp"
)

var ipCommand = &cobra.Command{
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
}
