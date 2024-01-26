package main

import (
	"flag"
	"log"
	"os"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/spf13/cobra"
	"wlb/cmd/job"
	"wlb/cmd/wlb"
	"wlb/internal/conf"
)

var rootCmd = &cobra.Command{
	Use:   "http",
	Short: "",
	Long:  ``,
}
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "./configs/config.yaml", "config path, eg: -conf config.yaml")

	rootCmd.AddCommand(wlb.HttpCmd)
	rootCmd.AddCommand(job.JobCmd)
}

func main() {
	flag.Parse()

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	job.Bc = bc
	wlb.Bc = bc

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
