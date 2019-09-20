package culture

import (
	"github.com/ironarachne/world/pkg/buildings"
	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/clothing"
	"github.com/ironarachne/world/pkg/food"
	"github.com/ironarachne/world/pkg/language"
	"github.com/ironarachne/world/pkg/music"
	"github.com/ironarachne/world/pkg/religion"
	"github.com/ironarachne/world/pkg/species"
)

// SimplifiedCulture is a simplified version of culture for display
type SimplifiedCulture struct {
	Adjective         string                            `json:"adjective"`
	AlcoholicDrinks   []string                          `json:"alcoholic_drinks"`
	BuildingStyle     buildings.SimplifiedBuildingStyle `json:"building_style"`
	ClothingStyle     clothing.SimplifiedStyle          `json:"clothing_style"`
	CommonFamilyNames []string                          `json:"common_family_names"`
	CommonFemaleNames []string                          `json:"common_female_names"`
	CommonMaleNames   []string                          `json:"common_male_names"`
	FoodStyle         food.SimplifiedStyle              `json:"food_style"`
	HomeClimate       climate.SimplifiedClimate         `json:"home_climate"`
	Language          language.SimplifiedLanguage       `json:"language"`
	MusicStyle        music.SimplifiedStyle             `json:"music_style"`
	Name              string                            `json:"name"`
	PrimaryRace       species.Simplified                `json:"primary_race"`
	Religion          religion.SimplifiedReligion       `json:"religion"`
	Views             []string                          `json:"views"`
}

// Simplify returns a simplified version of a culture for display
func (culture Culture) Simplify() SimplifiedCulture {
	drinks := []string{}
	for _, d := range culture.AlcoholicDrinks {
		drinks = append(drinks, d.Description)
	}
	return SimplifiedCulture{
		Adjective:         culture.Adjective,
		AlcoholicDrinks:   drinks,
		BuildingStyle:     culture.BuildingStyle.Simplify(),
		ClothingStyle:     culture.ClothingStyle.Simplify(),
		CommonFamilyNames: culture.CommonFamilyNames,
		CommonFemaleNames: culture.CommonFemaleNames,
		CommonMaleNames:   culture.CommonMaleNames,
		FoodStyle:         culture.FoodStyle.Simplify(),
		HomeClimate:       culture.HomeClimate.Simplify(),
		Language:          culture.Language.Simplify(),
		MusicStyle:        culture.MusicStyle.Simplify(),
		Name:              culture.Name,
		PrimaryRace:       culture.PrimaryRace.Simplify(),
		Religion:          culture.Religion.Simplify(),
		Views:             culture.Views,
	}
}
