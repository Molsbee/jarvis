package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var (
	say     bool
	rootCmd = &cobra.Command{
		Use:   "jarvis",
		Short: "Jarvis is an assistant that will perform a set of predefined functions",
		Long: `Jarvis is an assistant that will perform a set of predefined functions.
The application will be expanded to provide additional functionality required for my daily activities`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func init() {
	rootCmd.AddCommand(showConfig, clcCommands(), hackerNewsCommands())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
