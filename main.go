package main

//go:generate statik -src=./config -m

import (
	"fmt"
	"os"

	"github.com/MH4GF/github-nippou/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
