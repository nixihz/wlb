package job

import (
	"github.com/spf13/cobra"
	"wlb/internal/clients"
	"wlb/internal/data/po"
)

// initDataCmd represents the initData command
var initDataCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化数据",
	Long:  `初始化数据`,
	Run: func(cmd *cobra.Command, args []string) {
		dataClient := clients.NewSqlite3Client(Bc.Data)
		defer dataClient.Close()
		dataClient.LogMode(true)
		dataClient.CreateTable(&po.GoVersion{})
	},
}

func init() {
	JobCmd.AddCommand(initDataCmd)
}
