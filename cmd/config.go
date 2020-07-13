package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/Molsbee/jarvis/config"
	"github.com/howeyc/gopass"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var showConfig = &cobra.Command{
	Use:   "show-config",
	Short: "Inspect the config file being sourced by the program",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("failed to parse config file please create valid config file in home directory")
			}
		}()

		j, _ := json.MarshalIndent(config.GetConfig(), "", "  ")
		fmt.Println(string(j))
	},
}

var writeConfig = &cobra.Command{
	Use:   "write-config",
	Short: "Prompts the user for information required to run jarvis",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		zendeskUsername := promptUser("Zendesk Username: ", reader)
		zendeskPassword := promptUser("Zendesk Password: ", nil)
		domainUsername := promptUser("T3N Domain Username: ", reader)
		domainPassword := promptUser("T3N Domain Password: ", nil)

		config.Write(config.Config{
			Zendesk: config.ZendeskCredentials{
				Username: zendeskUsername,
				Password: zendeskPassword,
			},
			Domain: config.DomainCredentials{
				Name:     "T3N",
				Username: domainUsername,
				Password: domainPassword,
			},
		})
	},
}

func promptUser(prompt string, reader *bufio.Reader) string {
	fmt.Print(prompt)
	if reader != nil {
		result, _ := reader.ReadString('\n')
		result = strings.Replace(result, "\r", "", -1)
		result = strings.Replace(result, "\n", "", -1)
		return result
	}

	bytes, _ := gopass.GetPasswd()
	return string(bytes)
}
