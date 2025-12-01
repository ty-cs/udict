package cmd

import (
	"fmt"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var Version = "dev"

func getVersionString() string {
	// Version can be overridden at build time using -ldflags.
	// For example: go build -ldflags="-X 'github.com/ty-cs/udict/cmd.Version=v1.2.3'"
	if Version != "dev" {
		return Version
	}
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return Version
	}
	return info.Main.Version
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of udict",
	Long:  `All software has versions. This is udict's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("udict version: %s\n", getVersionString())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
