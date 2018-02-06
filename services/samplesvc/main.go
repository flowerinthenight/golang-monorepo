package main

import (
	goflag "flag"
	"fmt"

	"github.com/flowerinthenight/golang-monorepo/pkg/util"
	"github.com/golang/glog"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
)

var version = "?"

var (
	rootCmd = &cobra.Command{
		Use:   "samplesvc",
		Short: "A samplesvc for reference in ouchan monorepo.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			goflag.Parse()
		},
	}
)

func init() {
	rootCmd.AddCommand(RunCmd())
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
}

func RunCmd() *cobra.Command {
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Sample subcommand.",
		Long:  "Sample subcommand.",
		Run: func(cmd *cobra.Command, args []string) {
			glog.Infof("Hello from samplesvc (version: %v)", version)
			// sample use of pkg
			glog.Error(util.Err2(fmt.Errorf("test error")))
			u1 := uuid.NewV4()
			_ = u1
		},
	}

	return runCmd
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		glog.Fatalf("root cmd execute failed: %v", err)
	}
}
