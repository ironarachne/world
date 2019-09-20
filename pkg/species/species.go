package species

import (
	"strconv"

	"github.com/ironarachne/world/pkg/age"
	"github.com/ironarachne/world/pkg/dice"
	"github.com/ironarachne/world/pkg/measurement"
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/size"
	"github.com/ironarachne/world/pkg/trait"
)

// Species is a species of living thing
type Species struct {
	Name           string
	PluralName     string
	Adjective      string
	Commonality    int
	PossibleTraits []trait.Template // Traits that individuals of this species *might* have
	CommonTraits   []trait.Template // Traits that all members of this species share
	AgeCategories  []age.Category
	MinHumidity    int
	MaxHumidity    int
	MinTemperature int
	MaxTemperature int
	Resources      []resource.Resource // These are resources that can be derived from this species
	Tags           []string
}

// RandomHeight returns a random height appropriate for the given gender, age category, and size category
func RandomHeight(gender string, ageCategory age.Category, sizeCategory size.Category) string {
	var genderModifier int

	baseHeight := sizeCategory.BaseHeight

	if gender == "male" {
		genderModifier = ageCategory.MaleHeightModifier
	} else {
		genderModifier = ageCategory.FemaleHeightModifier
	}

	diceRoll := dice.Roll(ageCategory.HeightRangeDice)

	heightInInches := baseHeight + genderModifier + diceRoll

	height := measurement.ToString(heightInInches)

	return height
}

// RandomWeight returns a random weight appropriate for the given gender, age category, and size category
func RandomWeight(gender string, ageCategory age.Category, sizeCategory size.Category) string {
	var genderModifier int

	baseWeight := sizeCategory.BaseWeight

	if gender == "male" {
		genderModifier = ageCategory.MaleWeightModifier
	} else {
		genderModifier = ageCategory.FemaleWeightModifier
	}

	diceRoll := dice.Roll(ageCategory.WeightRangeDice)

	weightInPounds := baseWeight + genderModifier + diceRoll

	weight := strconv.Itoa(weightInPounds) + " lbs."

	return weight
}
