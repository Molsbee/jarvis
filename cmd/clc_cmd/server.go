package clc_cmd

import (
	"fmt"
	"github.com/Molsbee/jarvis/config"
	"github.com/Molsbee/jarvis/service/clc"
	"github.com/spf13/cobra"
	"os"
)

var (
	serverName   string
	hardwareUUID string
	environment  string
)

func vm() *cobra.Command {
	vm := &cobra.Command{
		Use:     "vm",
		Example: "vm --serverName UC1T3NMT2016-01 --env prod",
		Run: func(cmd *cobra.Command, args []string) {
			environment, err := config.GetEnvironment(environment)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// TODO: Could move from persistent flags and perform REGEX to determine if the server name is UUID.
			if len(serverName) != 0 {
				server, err := clc.GetServerDetails(environment, serverName)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				fmt.Println(server)
			} else if len(hardwareUUID) != 0 {
				server, err := clc.GetServerDetailsByHardwareUUID(environment, hardwareUUID)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				fmt.Println(server)
			} else {
				fmt.Println("please provide either a server name or hardware uuid")
				os.Exit(1)
			}
		},
	}
	vm.PersistentFlags().StringVarP(&serverName, "serverName", "s", "", "--serverName UC1T3NMT2016-01")
	vm.PersistentFlags().StringVarP(&hardwareUUID, "hardwareUUID", "u", "", "--hardwareUUID 8f295b9c7122487ab0b62c6a9b523ec2")
	vm.PersistentFlags().StringVarP(&environment, "env", "e", "prod", "--env dev (defaults to prod)")
	return vm
}
