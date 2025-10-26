package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of udict",
	Long:  `All software has versions. This is udict's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("udict v1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
