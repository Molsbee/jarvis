package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"time"
)

var moveReminder = &cobra.Command{
	Use:   "move-reminder",
	Short: "sets a reminder to prompt user to move periodically based on input (default 10m)",
	Run: func(cmd *cobra.Command, args []string) {
		duration, _ := time.ParseDuration("10m")
		if len(args) != 0 {
			var err error
			duration, err = time.ParseDuration(args[0])
			if err != nil {
				fmt.Println("please provide a valid duration for the timer (ex. 10m, 1h)")
				os.Exit(1)
			}
		}

		for {
			_ = <-time.NewTimer(duration).C
			exec.Command("say", "get up and move keyboard monkey").Run()
		}
	},
}
