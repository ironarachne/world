/*
Package culture provides methods and tools for generating fantasy cultures.
*/
package culture

import (
	"context"
	"fmt"

	"github.com/ironarachne/world/pkg/geography"
	"github.com/ironarachne/world/pkg/random"

	"github.com/ironarachne/world/pkg/buildings"
	"github.com/ironarachne/world/pkg/clothing"
	"github.com/ironarachne/world/pkg/conlang"
	"github.com/ironarachne/world/pkg/drink"
	"github.com/ironarachne/world/pkg/food"
	"github.com/ironarachne/world/pkg/language"
	"github.com/ironarachne/world/pkg/music"
	"github.com/ironarachne/world/pkg/race"
	"github.com/ironarachne/world/pkg/religion"
	"github.com/ironarachne/world/pkg/species"
)

const cultureError = "failed to generate culture: %w"

// Culture is a fantasy culture
type Culture struct {
	Adjective       string        `json:"adjective"`
	AlcoholicDrinks []drink.Drink `json:"alcoholic_drinks"`
	AttributeMax    int           `json:"attribute_max"`
	Attributes      struct {
		Aggression      int `json:"aggression"`
		Curiosity       int `json:"curiosity"`
		MagicPrevalence int `json:"magic_prevalence"`
		MagicStrength   int `json:"magic_strength"`
		Rigidity        int `json:"rigidity"`
		Superstition    int `json:"superstition"`
	} `json:"attributes"`
	BuildingStyle     buildings.BuildingStyle `json:"building_style"`
	ClothingStyle     clothing.Style          `json:"clothing_style"`
	CommonFamilyNames []string                `json:"common_family_names"`
	CommonFemaleNames []string                `json:"common_female_names"`
	CommonMaleNames   []string                `json:"common_male_names"`
	DrinkStyle        drink.Style             `json:"drink_style"`
	FoodStyle         food.Style              `json:"food_style"`
	Language          language.Language       `json:"language"`
	MusicStyle        music.Style             `json:"music_style"`
	Name              string                  `json:"name"`
	HomeBiome         string                  `json:"home_biome"`
	PrimaryRace       species.Species         `json:"primary_race"`
	Religion          religion.Religion       `json:"religion"`
	Views             []string                `json:"views"`
}

// Generate generates a culture
func Generate(ctx context.Context, home geography.Area) (Culture, error) {
	culture := Culture{}

	resources := home.GetResources()

	cultureLanguage, cultureLanguageCategory, err := conlang.Generate(ctx)
	if err != nil {
		err = fmt.Errorf(cultureError, err)
		return Culture{}, err
	}

	wordList, err := culture.createWordList(ctx, cultureLanguage, cultureLanguageCategory, home)
	if err != nil {
		err = fmt.Errorf(cultureError, err)
		return Culture{}, err
	}

	cultureLanguage.WordList = wordList
	culture.Language = cultureLanguage

	maleNames, err := culture.Language.RandomNameList(ctx, 10, "male")
	if err != nil {
		err = fmt.Errorf(cultureError, err)
		return Culture{}, err
	}
	culture.CommonMaleNames = maleNames
	femaleNames, err := culture.Language.RandomNameList(ctx, 10, "female")
	if err != nil {
		err = fmt.Errorf(cultureError, err)
		return Culture{}, err
	}
	culture.CommonFemaleNames = femaleNames
	familyNames, err := culture.Language.RandomNameList(ctx, 10, "family")
	if err != nil {
		err = fmt.Errorf(cultureError, err)
		return Culture{}, err
	}
	culture.CommonFamilyNames = familyNames

	culture.Name = culture.Language.Name
	culture.Adjective = culture.Language.Adjective

	musicStyle, err := music.Generate(ctx)
	if err != nil {
		err = fmt.Errorf(cultureError, err)
		return Culture{}, err
	}
	culture.MusicStyle = musicStyle

	buildingStyle, err := buildings.GenerateStyle(ctx)
	if err != nil {
		err = fmt.Errorf(cultureError, err)
		return Culture{}, err
	}
	culture.BuildingStyle = buildingStyle
	clothingStyle, err := clothing.GenerateStyle(ctx, home.Region.Temperature, resources)
	if err != nil {
		err = fmt.Errorf(cultureError, err)
		return Culture{}, err
	}
	culture.ClothingStyle = clothingStyle
	drinkStyle, err := drink.Generate(ctx, culture.Language, resources)
	if err != nil {
		err = fmt.Errorf(cultureError, err)
		return Culture{}, err
	}
	culture.DrinkStyle = drinkStyle
	foodStyle, err := food.GenerateStyle(ctx, resources)
	if err != nil {
		err = fmt.Errorf(cultureError, err)
		return Culture{}, err
	}
	culture.FoodStyle = foodStyle
	drinks, err := drink.RandomSet(ctx, 3, resources)
	if err != nil {
		err = fmt.Errorf(cultureError, err)
		return Culture{}, err
	}
	culture.AlcoholicDrinks = drinks

	culture.AttributeMax = 100
	culture.Attributes.Aggression = random.Intn(ctx, culture.AttributeMax) + 1
	culture.Attributes.Curiosity = random.Intn(ctx, culture.AttributeMax) + 1
	culture.Attributes.MagicPrevalence = random.Intn(ctx, culture.AttributeMax) + 1
	culture.Attributes.MagicStrength = random.Intn(ctx, culture.AttributeMax) + 1
	culture.Attributes.Rigidity = random.Intn(ctx, culture.AttributeMax) + 1
	culture.Attributes.Superstition = random.Intn(ctx, culture.AttributeMax) + 1

	primaryRace, err := race.RandomWeighted(ctx)
	if err != nil {
		err = fmt.Errorf(cultureError, err)
		return Culture{}, err
	}
	culture.PrimaryRace = primaryRace

	cultureReligion, err := religion.Generate(ctx, culture.Language)
	if err != nil {
		err = fmt.Errorf(cultureError, err)
		return Culture{}, err
	}
	culture.Religion = cultureReligion

	culture.Views = culture.getViews()

	culture.HomeBiome = home.Biome.Name

	return culture, nil
}

// Random returns a completely random culture
func Random(ctx context.Context) (Culture, error) {
	area, err := geography.Generate(ctx)
	if err != nil {
		err = fmt.Errorf(cultureError, err)
		return Culture{}, err
	}
	culture, err := Generate(ctx, area)
	if err != nil {
		err = fmt.Errorf(cultureError, err)
		return Culture{}, err
	}
	return culture, nil
}

func (culture Culture) createWordList(ctx context.Context, cultureLanguage language.Language, langCategory conlang.Category, area geography.Area) (map[string]string, error) {
	list := culture.Language.WordList

	for _, i := range area.Animals {
		if !conlang.IsInWordList(i.Name, list) {
			modifiedList, err := conlang.AddNounToWordList(ctx, cultureLanguage, langCategory, i.Name)
			if err != nil {
				err = fmt.Errorf(cultureError, err)
				return list, err
			}
			list = modifiedList
		}
	}

	for _, i := range area.Minerals {
		if !conlang.IsInWordList(i.Name, list) {
			modifiedList, err := conlang.AddNounToWordList(ctx, cultureLanguage, langCategory, i.Name)
			if err != nil {
				err = fmt.Errorf(cultureError, err)
				return list, err
			}
			list = modifiedList
		}
	}

	for _, i := range area.Plants {
		if !conlang.IsInWordList(i.Name, list) {
			modifiedList, err := conlang.AddNounToWordList(ctx, cultureLanguage, langCategory, i.Name)
			if err != nil {
				err = fmt.Errorf(cultureError, err)
				return list, err
			}
			list = modifiedList
		}
	}

	for _, i := range area.Seasons {
		if !conlang.IsInWordList(i.Name, list) {
			modifiedList, err := conlang.AddNounToWordList(ctx, cultureLanguage, langCategory, i.Name)
			if err != nil {
				err = fmt.Errorf(cultureError, err)
				return list, err
			}
			list = modifiedList
		}
	}

	return list, nil
}
