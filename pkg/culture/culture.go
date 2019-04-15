package culture

import (
	"math/rand"
)

// Culture is a fantasy culture
type Culture struct {
	Name              string
	Adjective         string
	Language          Language
	Appearance        Appearance
	CommonMaleNames   []string
	CommonFamilyNames []string
	CommonFemaleNames []string
	MusicStyle        MusicStyle
	AttributeMax      int
	Aggression        int
	Curiosity         int
	Rigidity          int
	Superstition      int
	HomeClimate       climate.Climate
	ClothingStyle     ClothingStyle
	FoodStyle         FoodStyle
	AlcoholicDrinks   []Drink
	Religion          Religion
}

// GenerateCulture generates a culture
func GenerateCulture() Culture {
	culture := Culture{}

	culture.Language = randomLanguage()

	culture.CommonMaleNames = culture.Language.GenerateNameList("male")
	culture.CommonFemaleNames = culture.Language.GenerateNameList("female")
	culture.CommonFamilyNames = culture.Language.GenerateNameList("family")

	culture.Name = culture.Language.Name
	culture.Adjective = culture.Language.Adjective

	culture.HomeClimate = climate.Generate()
	culture.MusicStyle = culture.randomMusicStyle()
	culture.ClothingStyle = culture.generateClothingStyle()
	culture.FoodStyle = culture.generateFoodStyle()
	culture.AlcoholicDrinks = culture.generateDrinks()

	culture.AttributeMax = 100
	culture.Aggression = rand.Intn(culture.AttributeMax) + 1
	culture.Curiosity = rand.Intn(culture.AttributeMax) + 1
	culture.Rigidity = rand.Intn(culture.AttributeMax) + 1
	culture.Superstition = rand.Intn(culture.AttributeMax) + 1

	culture.Appearance = culture.generateAppearance()

	culture.Religion = culture.generateReligion()

	return culture
}

// SetClimate sets the climate and recalculates some traits
func (culture Culture) SetClimate(climate string) Culture {
	newCulture := culture
	newCulture.HomeClimate = climate.GetClimate(climate)
	newCulture.MusicStyle = newCulture.randomMusicStyle()
	newCulture.ClothingStyle = newCulture.generateClothingStyle()
	newCulture.FoodStyle = newCulture.generateFoodStyle()
	newCulture.AlcoholicDrinks = newCulture.generateDrinks()

	return newCulture
}
