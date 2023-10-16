package cmd

import (
	"fmt"
	"os"

	"github.com/MH4GF/github-nippou/lib"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize github-nippou settings",
	Run: func(cmd *cobra.Command, args []string) {
		if err := lib.Init(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
