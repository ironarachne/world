/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ironarachne/world/config"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/region"
	"github.com/spf13/cobra"
)

// regionCmd represents the region command
var regionCmd = &cobra.Command{
	Use:   "region",
	Short: "Generate a random region",
	Long: `This command generates a random region.`,
	Run: func(cmd *cobra.Command, args []string) {
		configFile, _ := cmd.Flags().GetString("config")
		config.LoadConfig(configFile)
		randomSeed := random.SeedString()
		random.SeedFromString(randomSeed)
		
		result, err := region.Random()
		if err != nil {
			fmt.Println(fmt.Errorf("Failed to generate region: %w", err))
			os.Exit(1)
		}

		res, err := json.Marshal(result)
		if err != nil {
			fmt.Println(fmt.Errorf("Failed to generate region: %w", err))
			os.Exit(1)
		}
		fmt.Print(string(res))
	},
}

func init() {
	rootCmd.AddCommand(regionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// regionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// regionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
