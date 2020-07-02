package clc_cmd

import (
	"fmt"
	"github.com/Molsbee/jarvis/service"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

func zendesk() *cobra.Command {
	zendesk := &cobra.Command{
		Use:     "zendesk",
		Example: "jarvis clc zendesk",
		Run: func(cmd *cobra.Command, args []string) {
			tickets, err := service.GetZendeskTickets()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			for _, t := range tickets {
				fmt.Printf("%d %-15s %-100s\n", t.ID, t.GetGroupName(), t.ShortenedSubject())
			}

			if say {
				exec.Command("say", fmt.Sprintf("Zendesk currently has %d tickets sitting in queue", len(tickets))).Run()
			}
		},
	}
	zendesk.PersistentFlags().BoolVarP(&say, "say", "s", false, "say (adds os call to run say if its installed)")

	return zendesk
}
