package job

import (
	"os"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/spf13/cobra"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string

	id, _ = os.Hostname()
)

// goReleaseVersionCmd represents the goReleaseVersion command
var goReleaseVersionCmd = &cobra.Command{
	Use:   "go-new-version",
	Short: "监控go新版本发布",
	Long:  `监控go新版本发布`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := log.With(log.NewStdLogger(os.Stdout),
			"ts", log.DefaultTimestamp,
			"caller", log.DefaultCaller,
			"service.id", id,
			"service.name", Name,
			"service.version", Version,
			"trace.id", tracing.TraceID(),
			"span.id", tracing.SpanID(),
		)

		job, cancel, err := NewGoNewVersionJob(Bc.Data, logger)
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
