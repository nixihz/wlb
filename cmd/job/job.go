package job

import (
	"github.com/spf13/cobra"
	"wlb/internal/conf"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// flagconf is the config flag.
	flagconf string
	Bc       conf.Bootstrap
)

// JobCmd 脚本入口
// 执行方式 ./main job layout(或其他脚本)
var JobCmd = &cobra.Command{
	Use: "job",
}
