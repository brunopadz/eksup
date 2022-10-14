package cmd

import (
	"os"
	"ueks/pkg/config"

	"github.com/spf13/cobra"
)

// initCmd represents the configure command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	RunE:  runInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func runInit(cmd *cobra.Command, _ []string) error {
	home, _ := os.UserHomeDir()
	file := "/.ueks.yaml"
	path := home + file

	create, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer create.Close()

	_, err = create.WriteString(config.SampleCfg)
	if err != nil {
		return err
	}

	return err
}
