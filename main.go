package main

import (
	"log"

	"github.com/spf13/cobra"
	"wlb/cmd/wlb"
)

var rootCmd = &cobra.Command{
	Use:   "http",
	Short: "",
	Long:  ``,
}

func init() {
	rootCmd.AddCommand(wlb.HttpCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
