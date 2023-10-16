package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/MH4GF/github-nippou/lib"
	"github.com/spf13/cobra"
)

var sinceDate string
var untilDate string
var debug bool
var auth lib.Auth

// RootCmd defines a root command
var RootCmd = &cobra.Command{
	Use:   "github-nippou",
	Short: "Print today's your GitHub action",
	Run: func(cmd *cobra.Command, args []string) {
		if err := auth.Init(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		lines, err := lib.List(sinceDate, untilDate, debug, auth)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(lines)
	},
}

func init() {
	nowDate := time.Now().Format("20060102")
	sinceDate = nowDate
	untilDate = nowDate
	auth = lib.Auth{}

	RootCmd.PersistentFlags().StringVarP(&sinceDate, "since-date", "s", sinceDate, "Retrieves GitHub user_events since the date")
	RootCmd.PersistentFlags().StringVarP(&untilDate, "until-date", "u", untilDate, "Retrieves GitHub user_events until the date")
	RootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Debug mode")
}
