package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "openai-cli-tool",
	Short: "CLI tool to interact with OpenAI API",
	Long:  "A CLI tool written in Go to call OpenAI API for various tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the OpenAI CLI tool!")
	},
}

// Execute starts the CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

