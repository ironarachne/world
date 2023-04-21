/*
Copyright © 2022 Ben Overmyer <ben@overmyer.net>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ironarachne/world/config"

	"github.com/ironarachne/world/pkg/culture"
	"github.com/ironarachne/world/pkg/random"
	"github.com/spf13/cobra"
)

// cultureCmd represents the culture command
var cultureCmd = &cobra.Command{
	Use:   "culture",
	Short: "Generate a fantasy culture",
	Long:  `This command generates a random fantasy culture.`,
	Run: func(cmd *cobra.Command, args []string) {
		configFile, _ := cmd.Flags().GetString("config")
		config.LoadConfig(configFile)
		outputFile, _ := cmd.Flags().GetString("output")
		randomSeed := random.SeedString()
		random.SeedFromString(randomSeed)
		randomCulture, err := culture.Random()
		if err != nil {
			fmt.Println(fmt.Errorf("failed to generate culture: %w", err))
			os.Exit(1)
		}

		res, err := json.Marshal(randomCulture)
		if err != nil {
			fmt.Println(fmt.Errorf("failed to convert culture to json: %w", err))
			os.Exit(1)
		}
		if outputFile != "" {
			err = os.WriteFile(outputFile, res, 0644)
			if err != nil {
				fmt.Println(fmt.Errorf("failed to write culture to file: %w", err))
				os.Exit(1)
			}
		} else {
			fmt.Print(string(res))
		}
	},
}

func init() {
	rootCmd.AddCommand(cultureCmd)
}
