package clc_cmd

import (
	"fmt"
	"github.com/Molsbee/jarvis/config"
	"github.com/Molsbee/jarvis/service/clc"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var (
	environment string
)

func vm() *cobra.Command {
	vm := &cobra.Command{
		Use:     "vm",
		Example: "vm [server name or hardware uuid] --env prod",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("please provide either a server name or hardware uuid")
				os.Exit(1)
			}

			environment, err := config.GetEnvironment(environment)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			serverName := args[0]
			hardwareUUID, parseErr := uuid.Parse(serverName)
			if parseErr == nil {
				nakedUUID := strings.ToLower(fmt.Sprintf("%X%X%X%X%X", hardwareUUID[0:4], hardwareUUID[4:6], hardwareUUID[6:8], hardwareUUID[8:10], hardwareUUID[10:]))
				server, err := clc.GetServerDetailsByHardwareUUID(environment, nakedUUID)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				fmt.Println(server)
			} else {
				server, err := clc.GetServerDetails(environment, serverName)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				fmt.Println(server)
			}
		},
	}

	vm.PersistentFlags().StringVarP(&environment, "env", "e", "prod", "--env dev (defaults to prod)")
	return vm
}
