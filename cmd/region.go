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
		outputFile, _ := cmd.Flags().GetString("output")
		randomSeed := random.SeedString()
		random.SeedFromString(randomSeed)
		
		result, err := region.Random()
		if err != nil {
			fmt.Println(fmt.Errorf("failed to generate region: %w", err))
			os.Exit(1)
		}

		res, err := json.Marshal(result)
		if err != nil {
			fmt.Println(fmt.Errorf("failed to generate region: %w", err))
			os.Exit(1)
		}
		if outputFile != "" {
			err = os.WriteFile(outputFile, res, 0644)
			if err != nil {
				fmt.Println(fmt.Errorf("failed to write region to file: %w", err))
				os.Exit(1)
			}
		} else {
			fmt.Print(string(res))
		}
	},
}

func init() {
	rootCmd.AddCommand(regionCmd)
}
