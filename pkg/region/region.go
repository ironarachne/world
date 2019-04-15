package region

import (
	"math/rand"
	"strings"

	"github.com/ironarachne/random"
	"github.com/ironarachne/world/pkg/character"
	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/culture"
	"github.com/ironarachne/world/pkg/heraldry"
	"github.com/ironarachne/world/pkg/organization"
	"github.com/ironarachne/world/pkg/town"
)

// Region is a map region
type Region struct {
	Biome         string
	Culture       culture.Culture
	Climate       climate.Climate
	Capital       string
	Class         RegionClass
	Name          string
	Ruler         character.Character
	RulerBlazon   string
	RulerHeraldry string
	RulerTitle    string
	Towns         []town.Town
	Organizations []organization.Organization
}

// RegionClass is a class of region
type RegionClass struct {
	MaxNumberOfTowns int
	MinNumberOfTowns int
	Name             string
	RulerTitleFemale string
	RulerTitleMale   string
}

func randomClass() RegionClass {
	class := random.ItemFromThresholdMap(classes)
	regionClass := classData[class]

	return regionClass
}

// Generate generates a random region
func Generate(regionType string) Region {
	region := Region{}
	biome := climate.Climate{}

	if regionType == "random" {
		biome = climate.Generate()
	} else {
		biome = climate.GetClimate(regionType)
	}

	regionType = biome.Name

	region.Biome = biome.Name
	region.Climate = biome
	region.Culture = culture.Generate()
	region.Culture = region.Culture.SetClimate(region.Biome)

	region.Class = randomClass()

	newTown := town.Generate("city", regionType)
	newTown = town.SetCulture(region.Culture, newTown)
	region.Towns = append(region.Towns, newTown)

	region.Capital = newTown.Name

	for i := region.Class.MinNumberOfTowns - 1; i < region.Class.MaxNumberOfTowns-1; i++ {
		newTown = town.Generate("random", regionType)
		newTown = town.SetCulture(region.Culture, newTown)
		region.Towns = append(region.Towns, newTown)
	}

	numberOfOrgs := rand.Intn(3) + 1
	newOrg := organization.Organization{}

	for i := 0; i < numberOfOrgs; i++ {
		newOrg = organization.Generate()
		region.Organizations = append(region.Organizations, newOrg)
	}

	region.Ruler = region.generateRuler()

	region.RulerTitle = region.Class.RulerTitleFemale
	if region.Ruler.Gender == "male" {
		region.RulerTitle = region.Class.RulerTitleMale
	}

	device := heraldry.Generate()
	region.RulerHeraldry = device.RenderToSVG(320, 420)
	region.RulerBlazon = device.RenderToBlazon()

	regionName := region.Culture.Language.RandomName()
	region.Name = strings.Title(regionName)

	return region
}
