/*
Copyright © 2022 Ben Overmyer <ben@overmyer.net>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/ironarachne/world/config"
	"os"

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
		randomSeed := random.SeedString()
		random.SeedFromString(randomSeed)
		randomCulture, err := culture.Random()
		if err != nil {
			fmt.Println(fmt.Errorf("Failed to generate culture: %w", err))
			os.Exit(1)
		}

		res, err := json.Marshal(randomCulture)
		if err != nil {
			fmt.Println(fmt.Errorf("Failed to generate culture: %w", err))
			os.Exit(1)
		}
		fmt.Print(string(res))
	},
}

func init() {
	rootCmd.AddCommand(cultureCmd)
	cultureCmd.PersistentFlags().StringP("config", "c", "", "Config file name.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cultureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cultureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
