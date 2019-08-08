package cmd

import (
	"fmt"
	"os"

	"github.com/wade-liwei/demo/rest"

	"github.com/spf13/cobra"
	//"github.com/spf13/viper"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "run main.go",
	Short: "read mysal table data.",

	Run: func(cmd *cobra.Command, args []string) {
		rest.WebServer()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
