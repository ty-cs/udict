package cmd

import (
	"fmt"
	"runtime/debug"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of udict",
	Long:  `All software has versions. This is udict's`,
	Run: func(cmd *cobra.Command, args []string) {
		info, ok := debug.ReadBuildInfo()
		if !ok {
			fmt.Println("udict version unknown (built without Go modules)")
			return
		}

		var (
			version = info.Main.Version
		)
		// print version details
		fmt.Printf("udict version: %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
