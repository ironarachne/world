package culture

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/clothing"
	"github.com/ironarachne/world/pkg/food"
	"github.com/ironarachne/world/pkg/language"
	"github.com/ironarachne/world/pkg/music"
	"github.com/ironarachne/world/pkg/race"
	"github.com/ironarachne/world/pkg/religion"
)

// Culture is a fantasy culture
type Culture struct {
	Name              string
	Adjective         string
	Language          language.Language
	PrimaryRace        race.Race
	CommonMaleNames   []string
	CommonFamilyNames []string
	CommonFemaleNames []string
	MusicStyle        music.Style
	AttributeMax      int
	Aggression        int
	Curiosity         int
	Rigidity          int
	Superstition      int
	HomeClimate       climate.Climate
	ClothingStyle     clothing.Style
	FoodStyle         food.Style
	AlcoholicDrinks   []food.Drink
	Religion          religion.Religion
}

// Generate generates a culture
func Generate(homeClimate climate.Climate) Culture {
	culture := Culture{}
	culture.HomeClimate = homeClimate

	culture.Language = language.Generate()

	culture.CommonMaleNames = culture.Language.GenerateNameList("male")
	culture.CommonFemaleNames = culture.Language.GenerateNameList("female")
	culture.CommonFamilyNames = culture.Language.GenerateNameList("family")

	culture.Name = culture.Language.Name
	culture.Adjective = culture.Language.Adjective

	instruments := music.GenerateInstruments(culture.HomeClimate)
	culture.MusicStyle = music.GenerateStyle(instruments)

	culture.ClothingStyle = clothing.GenerateStyle(culture.HomeClimate)
	culture.FoodStyle = food.GenerateStyle(culture.HomeClimate)
	culture.AlcoholicDrinks = food.GenerateDrinks(culture.HomeClimate)

	culture.AttributeMax = 100
	culture.Aggression = rand.Intn(culture.AttributeMax) + 1
	culture.Curiosity = rand.Intn(culture.AttributeMax) + 1
	culture.Rigidity = rand.Intn(culture.AttributeMax) + 1
	culture.Superstition = rand.Intn(culture.AttributeMax) + 1

	culture.PrimaryRace = race.GetRandom()

	culture.Religion = religion.Generate(culture.Language)

	return culture
}

// Random returns a completely random culture
func Random() Culture {
	homeClimate := climate.Generate()
	return Generate(homeClimate)
}
