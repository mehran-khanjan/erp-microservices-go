package cmd

import (
	"fmt"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "help",
	Short: "Help command",
	Long:  "Display help command",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
