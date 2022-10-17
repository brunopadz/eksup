package cmd

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "check",
	Short: "A brief description of your command",
}

func init() {
	rootCmd.AddCommand(listCmd)
}
