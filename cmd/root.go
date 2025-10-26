package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "udict",
	Short: "A CLI tool to get definitions from Urban Dictionary",
	Long: `udict is a command-line tool that allows you to quickly get definitions
from Urban Dictionary. You can use the 'define' subcommand to look up a word.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
