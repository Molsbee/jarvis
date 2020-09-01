package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"time"
)

var stopWatch = &cobra.Command{
	Use:   "stop-watch",
	Short: "start a stop watch for a provided duration that is longer than 1 second (ex. 30s, 1m)",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("please provide a valid duration for the stop watch (ex. 300s, 1m) that is greater than 1 second")
			os.Exit(1)
		}

		duration, err := time.ParseDuration(args[0])
		if err != nil || duration.Seconds() < 1 {
			fmt.Println("please provide a valid duration for the stop watch (ex. 300s, 1m) that is greater that 1 second")
			os.Exit(1)
		}

		for i := 0; i < int(duration.Seconds()); i++ {
			if i != 0 && i%60 == 0 {
				fmt.Println()
			}

			fmt.Print("|")
			time.Sleep(1 * time.Second)
		}
		fmt.Printf("\nTimer has completed\n")
		exec.Command("say", "beep beep beep beep beep beep").Run()
	},
}
