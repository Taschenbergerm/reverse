package cmd

import (
	"fmt"
	"os"

	cobra "github.com/spf13/cobra"
)

var cfgFike string

var rootCmd = &cobra.Command{
	Use: "reverse <command> [flags]",
	Short: "reverse is a go utility to reverse a string",
	Long: "reverse can reverse string from various source and write to various sources"}

// Execute add all child commands to the root command and set flags appropriately,
// This is called by main.main. It only needs to happen once to the rootCmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init(){
	cobra.OnInitialize()
}