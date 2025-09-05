package main

import (
	"context"
	goflag "flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/flowerinthenight/golang-monorepo/internal"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
)

var version = "?"

var (
	rootCmd = &cobra.Command{
		Use:   "samplesvc",
		Short: "samplesvc for reference",
		Long:  "A samplesvc for reference in golang-monorepo.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			goflag.Parse()
		},
	}
)

func init() {
	rootCmd.AddCommand(RunCmd())
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
}

func run(quit context.Context, done chan error) {
	// Sample use of internal package.
	glog.Error(internal.GetHostname())

	<-quit.Done()
	done <- nil
}

func RunCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "run samplesvc",
		Long:  "Run samplesvc as a long-running service.",
		RunE: func(cmd *cobra.Command, args []string) error {
			defer func(begin time.Time) {
				glog.Infof("stop samplesvc after %v", time.Since(begin))
			}(time.Now())

			glog.Infof("start samplesvc on %v", time.Now())

			quit, cancel := context.WithCancel(context.TODO())
			done := make(chan error)
			go run(quit, done)

			go func() {
				sigch := make(chan os.Signal, 1)
				signal.Notify(sigch, syscall.SIGINT, syscall.SIGTERM)
				glog.Info(<-sigch)
				cancel()
			}()

			return <-done
		},
	}

	return cmd
}

func main() {
	rootCmd.Execute()
}
