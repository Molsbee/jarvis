package clc_cmd

import (
	"github.com/spf13/cobra"
)

var say bool

func GetCommands() *cobra.Command {
	clcCommand := &cobra.Command{
		Use:   "clc",
		Short: "Helpful commands for interacting with Century Link Cloud",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	checkMK.PersistentFlags().BoolVarP(&say, "say", "s", false, "say (adds os call to run say if its installed)")
	clcCommand.AddCommand(ipCommand, vm(), haProxy(), zendesk(), checkMK)
	return clcCommand
}
