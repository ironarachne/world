package culture

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/buildings"
	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/clothing"
	"github.com/ironarachne/world/pkg/drink"
	"github.com/ironarachne/world/pkg/food"
	"github.com/ironarachne/world/pkg/language"
	"github.com/ironarachne/world/pkg/music"
	"github.com/ironarachne/world/pkg/race"
	"github.com/ironarachne/world/pkg/religion"
)

// Culture is a fantasy culture
type Culture struct {
	Adjective       string
	AlcoholicDrinks []drink.Drink
	AttributeMax    int
	Attributes      struct {
		Aggression      int
		Curiosity       int
		MagicPrevalence int
		MagicStrength   int
		Rigidity        int
		Superstition    int
	}
	BuildingStyle     buildings.BuildingStyle
	ClothingStyle     clothing.Style
	CommonFamilyNames []string
	CommonFemaleNames []string
	CommonMaleNames   []string
	FoodStyle         food.Style
	HomeClimate       climate.Climate
	Language          language.Language
	MusicStyle        music.Style
	Name              string
	PrimaryRace       race.Race
	Religion          religion.Religion
	Views             []string
}

// Generate generates a culture
func Generate(homeClimate climate.Climate) Culture {
	culture := Culture{}
	culture.HomeClimate = homeClimate

	culture.Language = language.Generate()

	for _, a := range homeClimate.Animals {
		if !language.IsInWordList(a.Name, culture.Language.WordList) {
			culture.Language.WordList = culture.Language.AddNounToWordList(a.Name)
		}
	}

	for _, p := range homeClimate.Plants {
		if !language.IsInWordList(p.Name, culture.Language.WordList) {
			culture.Language.WordList = culture.Language.AddNounToWordList(p.Name)
		}
	}

	for _, f := range homeClimate.Fish {
		if !language.IsInWordList(f.Name, culture.Language.WordList) {
			culture.Language.WordList = culture.Language.AddNounToWordList(f.Name)
		}
	}

	culture.CommonMaleNames = culture.Language.GenerateNameList("male")
	culture.CommonFemaleNames = culture.Language.GenerateNameList("female")
	culture.CommonFamilyNames = culture.Language.GenerateNameList("family")

	culture.Name = culture.Language.Name
	culture.Adjective = culture.Language.Adjective

	instruments := music.GenerateInstruments(culture.HomeClimate)
	culture.MusicStyle = music.GenerateStyle(instruments)

	culture.BuildingStyle = buildings.GenerateStyle()
	culture.ClothingStyle = clothing.GenerateStyle(culture.HomeClimate)
	culture.FoodStyle = food.GenerateStyle(culture.HomeClimate)
	culture.AlcoholicDrinks = drink.RandomSet(3, culture.HomeClimate.Resources)

	culture.AttributeMax = 100
	culture.Attributes.Aggression = rand.Intn(culture.AttributeMax) + 1
	culture.Attributes.Curiosity = rand.Intn(culture.AttributeMax) + 1
	culture.Attributes.MagicPrevalence = rand.Intn(culture.AttributeMax) + 1
	culture.Attributes.MagicStrength = rand.Intn(culture.AttributeMax) + 1
	culture.Attributes.Rigidity = rand.Intn(culture.AttributeMax) + 1
	culture.Attributes.Superstition = rand.Intn(culture.AttributeMax) + 1

	parentRace := race.GetRandom()
	culture.PrimaryRace = race.GenerateSubrace(parentRace)

	culture.Religion = religion.Generate(culture.Language)

	culture.Views = culture.getViews()

	return culture
}

// Random returns a completely random culture
func Random() Culture {
	homeClimate := climate.Generate()
	return Generate(homeClimate)
}
