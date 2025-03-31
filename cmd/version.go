package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/waktusolatmy/cli/common"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version will output the current build information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("CLI version: v%s\n", common.VersionCli)
		fmt.Printf("Go version: %s\n", runtime.Version())
		fmt.Printf("Build date: %s\n", common.DateCli)
		fmt.Printf("Git commit: %s\n", common.CommitCli)
		fmt.Printf("OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
