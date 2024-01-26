package job

import (
	"github.com/spf13/cobra"
)

// goReleaseVersionCmd represents the goReleaseVersion command
var goReleaseVersionCmd = &cobra.Command{
	Use:   "go-new-version",
	Short: "监控go新版本发布",
	Long:  `监控go新版本发布`,
	Run: func(cmd *cobra.Command, args []string) {
		job, cancel, err := NewGoNewVersionJob(Bc.Data)
		if err != nil {
			panic(err)
		}
		defer cancel()
		job.Run()
	},
}

func init() {
	JobCmd.AddCommand(goReleaseVersionCmd)
}
