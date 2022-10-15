package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"ueks/pkg/config"
)

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
	fmt.Printf("Generating ueks config file at %s\n", path)

	_, err = create.WriteString(config.SampleCfg)
	if err != nil {
		return err
	}

	fmt.Println("Config file created. Edit it configure AWS authentication method.")
	return nil
}
