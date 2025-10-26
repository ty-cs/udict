package cmd

import (
	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random word definition from Urban Dictionary",
	Long: `This command fetches a random word definition from Urban Dictionary
and displays it in a styled format.`,
	Run: func(cmd *cobra.Command, args []string) {
		getAndDisplayDefinition("random", 1, "No random definition found 😕")
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}
