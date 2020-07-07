package clc_cmd

import (
	"fmt"
	"github.com/Molsbee/jarvis/service"
	"github.com/spf13/cobra"
	"os/exec"
)

var checkMK = &cobra.Command{
	Use: "check_mk",
	Run: func(cmd *cobra.Command, args []string) {
		alerts, err := service.GetCheckMKStatus()
		if err != nil {
			fmt.Println(err)
		}

		for _, alert := range alerts {
			fmt.Println(alert)
		}
		if say {
			exec.Command("say", fmt.Sprintf("CheckMK currently has %d errors slash warnings displayed", len(alerts))).Run()
		}
	},
}
