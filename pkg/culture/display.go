package culture

import (
	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/clothing"
	"github.com/ironarachne/world/pkg/food"
	"github.com/ironarachne/world/pkg/language"
	"github.com/ironarachne/world/pkg/music"
	"github.com/ironarachne/world/pkg/race"
	"github.com/ironarachne/world/pkg/religion"
)

// SimplifiedCulture is a simplified version of culture for display
type SimplifiedCulture struct {
	Adjective         string                      `json:"adjective"`
	AlcoholicDrinks   []food.Drink                `json:"alcoholic_drinks"`
	ClothingStyle     clothing.Style              `json:"clothing_style"`
	CommonFamilyNames []string                    `json:"common_family_names"`
	CommonFemaleNames []string                    `json:"common_female_names"`
	CommonMaleNames   []string                    `json:"common_male_names"`
	FoodStyle         food.Style                  `json:"food_style"`
	HomeClimate       climate.SimplifiedClimate   `json:"home_climate"`
	Language          language.SimplifiedLanguage `json:"language"`
	MusicStyle        music.SimplifiedStyle       `json:"music_style"`
	Name              string                      `json:"name"`
	PrimaryRace       race.SimplifiedRace         `json:"primary_race"`
	Religion          religion.SimplifiedReligion `json:"religion"`
}

// Simplify returns a simplified version of a culture for display
func (culture Culture) Simplify() SimplifiedCulture {
	return SimplifiedCulture{
		Adjective:         culture.Adjective,
		AlcoholicDrinks:   culture.AlcoholicDrinks,
		ClothingStyle:     culture.ClothingStyle,
		CommonFamilyNames: culture.CommonFamilyNames,
		CommonFemaleNames: culture.CommonFemaleNames,
		CommonMaleNames:   culture.CommonMaleNames,
		FoodStyle:         culture.FoodStyle,
		HomeClimate:       culture.HomeClimate.Simplify(),
		Language:          culture.Language.Simplify(),
		MusicStyle:        culture.MusicStyle.Simplify(),
		Name:              culture.Name,
		PrimaryRace:       culture.PrimaryRace.Simplify(),
		Religion:          culture.Religion.Simplify(),
	}
}
