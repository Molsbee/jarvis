package cmd

import (
	"fmt"
	"github.com/Molsbee/jarvis/service/hacker_news"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func hackerNewsCommands() *cobra.Command {
	hackerNews := &cobra.Command{
		Use:   "hacker-news",
		Short: "Commands for interacting with hacker news.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	hackerNews.AddCommand(&cobra.Command{
		Use:   "top",
		Short: "List the top stories on hacker news",
		Long:  "List the top stories on hacker news (optionally pass in a number for how many stories to show defaults to 20)",
		Run: func(cmd *cobra.Command, args []string) {
			value := parseStoryCount(args, 20)
			stories, err := hacker_news.GetTopStories(value)
			if err != nil {
				fmt.Printf("unable to load the top stories from hacker news - %s", err)
				os.Exit(1)
			}

			for i, story := range stories {
				fmt.Printf("%-4d%-65s%s\n", i+1, story.GetTitle(), story.URL)
			}
		},
	},
		&cobra.Command{
			Use:   "new",
			Short: "List new stories on hacker news",
			Long:  "List new stories on hacker news (optionally pass in a number for how many stories to show defaults to 20)",
			Run: func(cmd *cobra.Command, args []string) {
				value := parseStoryCount(args, 20)
				stories, err := hacker_news.GetNewStories(value)
				if err != nil {
					fmt.Printf("unable to load the newest stories from hacker news - %s", err)
					os.Exit(1)
				}

				for i, story := range stories {
					fmt.Printf("%-4d%-65s%s\n", i+1, story.GetTitle(), story.URL)
				}
			},
		},
		&cobra.Command{
			Use:   "best",
			Short: "List the best stories on hacker news",
			Long:  "List the best stories on hacker news (optionally pass in a number for how many stories to show defaults to 20)",
			Run: func(cmd *cobra.Command, args []string) {
				value := parseStoryCount(args, 20)
				stories, err := hacker_news.GetBestStories(value)
				if err != nil {
					fmt.Printf("unable to load the best stories from hacker news - %s", err)
					os.Exit(1)
				}

				for i, story := range stories {
					fmt.Printf("%-4d%-65s%s\n", i+1, story.GetTitle(), story.URL)
				}
			},
		})

	return hackerNews
}

func parseStoryCount(args []string, defaultValue int) int {
	if len(args) != 0 {
		v, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Please provide a valid number")
			os.Exit(1)
		}
		return v
	}

	return defaultValue
}
