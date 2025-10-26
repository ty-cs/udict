package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var limit int

const maxDefinitions = 5

// defineCmd represents the define command
var defineCmd = &cobra.Command{
	Use:   "define [word]",
	Short: "Get the definition of a word from Urban Dictionary",
	Long: `This command fetches the definition of a word from Urban Dictionary
and displays it in a styled format.`,
	Args: cobra.ExactArgs(1),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if limit > maxDefinitions {
			limit = maxDefinitions
		}
		if limit < 1 {
			return fmt.Errorf("limit cannot be less than 1")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		word := args[0]
		if strings.TrimSpace(word) == "" {
			fmt.Println("Word cannot be empty.")
			return
		}
		endpoint := "define?term=" + word
		notFoundMsg := fmt.Sprintf("No definition found for %s 😕\n", word)
		getAndDisplayDefinition(endpoint, limit, notFoundMsg)
	},
}

func init() {
	rootCmd.AddCommand(defineCmd)
	defineCmd.Flags().IntVarP(&limit, "limit", "l", 1, "Number of definitions to show (max 5)")
}
