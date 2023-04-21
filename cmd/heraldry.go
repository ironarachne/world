/*
Copyright © 2022 Ben Overmyer <ben@overmyer.net>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ironarachne/world/config"

	"github.com/ironarachne/world/pkg/heraldry"
	"github.com/ironarachne/world/pkg/random"
	"github.com/spf13/cobra"
)

var fieldName string
var chargeTag string

// heraldryCmd represents the heraldry command
var heraldryCmd = &cobra.Command{
	Use:   "heraldry",
	Short: "Generate fantasy heraldry",
	Long:  `This command generates fantasy heraldry.`,
	Run: func(cmd *cobra.Command, args []string) {
		configFile, _ := cmd.Flags().GetString("config")
		config.LoadConfig(configFile)
		outputFile, _ := cmd.Flags().GetString("output")
		randomSeed := random.SeedString()
		random.SeedFromString(randomSeed)
		fieldType := fieldName

		var o heraldry.Device
		var err error

		o, err = heraldry.GenerateByParameters(fieldType, chargeTag)
		if err != nil {
			fmt.Println(fmt.Errorf("failed to generate heraldry: %w", err))
			os.Exit(1)
		}

		sd := o.Simplify()

		res, err := json.Marshal(sd)
		if err != nil {
			fmt.Println(fmt.Errorf("failed to convert heraldry to json: %w", err))
			os.Exit(1)
		}
		if outputFile != "" {
			err = os.WriteFile(outputFile, res, 0644)
			if err != nil {
				fmt.Println(fmt.Errorf("failed to write heraldry to file: %w", err))
				os.Exit(1)
			}
		} else {
			fmt.Print(string(res))
		}
	},
}

func init() {
	rootCmd.AddCommand(heraldryCmd)
	heraldryCmd.Flags().StringVarP(&fieldName, "field", "f", "", "Name of the field to use")
	heraldryCmd.Flags().StringVarP(&chargeTag, "tag", "t", "", "Tag to use for picking charges")
}
