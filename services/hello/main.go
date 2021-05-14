package main

import (
	"context"
	goflag "flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chibimi/golang-monorepo/pkg/util"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
)

var version = "?"

var (
	rootCmd = &cobra.Command{
		Use:   "hello",
		Short: "hello for reference",
		Long:  "A hello for reference in golang-monorepo.",
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
	// sample use of pkg
	glog.Error(util.Err2(fmt.Errorf("test error")))

	for {
		select {
		case <-quit.Done():
			done <- nil
			return
		}
	}
}

func RunCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "run hello",
		Long:  "Run hello as a long-running service.",
		RunE: func(cmd *cobra.Command, args []string) error {
			defer func(begin time.Time) {
				glog.Infof("stop hello after %v", time.Since(begin))
			}(time.Now())

			glog.Infof("start hello on %v", time.Now())

			quit, cancel := context.WithCancel(context.TODO())
			done := make(chan error)
			go run(quit, done)

			go func() {
				sigch := make(chan os.Signal)
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
	if err := rootCmd.Execute(); err != nil {
		glog.Fatalf("%v", err)
	}
}
