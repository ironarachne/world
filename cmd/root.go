/*
Copyright © 2023 Ben Overmyer <ben@overmyer.net>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Version = "0.0.0"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "world",
	Short: "A tool for generating fictional worlds",
	Long: `This is a program for generating fictional worlds.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "", "Config file name.")
	rootCmd.PersistentFlags().StringP("output", "o", "", "File name to write to. Default is stdout.")
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of world",
	Long:  `Print the version number of world in semantic versioning format.`,
	Run: func(cmd *cobra.Command, args []string) {
	  fmt.Println("world v" + Version)
	},
  }
